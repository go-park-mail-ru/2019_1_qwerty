package helpers

import (
	"2019_1_qwerty/database"
	"2019_1_qwerty/models"
	"math"
	"math/rand"
	"sync"
	"time"
)

//PlayerState - player state in room->game
type PlayerState struct {
	ID    string
	X     int
	Y     int
	Score int
}

//ObjectState - object state in room->game
type ObjectState struct {
	ID    int
	X     int
	Y     int
	Speed int
}

//RoomState - room model state
type RoomState struct {
	Players     map[string]PlayerState
	Objects     []ObjectState
	CurrentTime time.Time
}

//Room - room model
type Room struct {
	ID         string
	MaxPlayers int
	Players    map[string]*Player
	mu         *sync.Mutex
	register   chan *Player
	unregister chan *Player
	ticker     *time.Ticker
	workTicker *time.Ticker
	state      *RoomState
}

//NewRoom - create new room
func NewRoom(maxPlayers int, id string) *Room {
	return &Room{
		ID:         id,
		MaxPlayers: maxPlayers,
		register:   make(chan *Player),
		unregister: make(chan *Player),
		ticker:     time.NewTicker(50 * time.Millisecond),   // HAS TO BEE 100 * ...
		workTicker: time.NewTicker(750 * time.Millisecond), // meteors
		Players:    make(map[string]*Player),
		mu:         new(sync.Mutex),
	}
}

//AddPlayer - add player to slice of players
func (r *Room) AddPlayer(player *Player) {
	r.register <- player
}

//RemovePlayer - removes player
func (r *Room) RemovePlayer(player *Player) {
	r.unregister <- player
}

//CreatePlayerState - initialize room at the start of the game
func CreatePlayerState(players map[string]*Player) *RoomState {
	state := &RoomState{
		Players:     make(map[string]PlayerState),
		Objects:     make([]ObjectState, 0),
		CurrentTime: time.Now(),
	}

	i := 0

	for _, player := range players {
		tmp := state.Players[player.number]
		tmp.ID = player.ID
		tmp.Score = 0
		tmp.X = 50
		tmp.Y = 50 * (i + 1)
		state.Players[player.number] = tmp
		i++
	}

	return state
}

//CreateObject - create meteor
func CreateObject(r *Room) []ObjectState {
	rand.Seed(time.Now().UnixNano())
	object := ObjectState{X: 350, Y: rand.Intn(135-5) + 5, Speed: rand.Intn(9-6) + 6}
	r.state.Objects = append(r.state.Objects, object)
	return r.state.Objects
}

const sqlInsertPlayer = `
	INSERT INTO scores(player, score) 
	VALUES($1, $2) 
`

const sqlUpdatePlayer = `
	UPDATE scores 
	SET score = GREATEST($2, (SELECT score FROM scores WHERE player = $1))
	WHERE player = $1
`

func insertScoreToDB(players map[string]*Player) {
	for _, player := range players {
		var name string
		err := database.Database.QueryRow("SELECT player FROM scores WHERE player = $1", player.ID).Scan(&name)

		if err != nil {
			_, _ = database.Database.Exec(sqlInsertPlayer, player.ID, player.score)
		} else {
			_, _ = database.Database.Exec(sqlUpdatePlayer, player.ID, player.score)
		}

	}
}

//Run - runs rooms
func (r *Room) Run() {
	for {
		select {
		case <-r.unregister:
			r.mu.Lock()
			for _, p := range r.Players {
				delete(r.Players, p.number)
				p.SendMessage(&models.Logs{Head: "GAME ENDED", Content: nil})
			}
			r.mu.Unlock()

			MainGame.mu.Lock()
			MainGame.RemoveRoom(r)
			MainGame.mu.Unlock()

			return

		case player := <-r.register:
			r.mu.Lock()
			r.Players[player.number] = player
			player.room = r
			r.mu.Unlock()
			player.SendMessage(&models.Logs{Head: "CONNECTED", Content: nil})

			if len(r.Players) == r.MaxPlayers {
				r.state = CreatePlayerState(r.Players)

				for _, player := range r.Players {
					coord := r.state.Players
					player.SendMessage(&models.Logs{Head: "GAME STARTED", Content: coord})
				}
			}

		case <-r.workTicker.C:
			if len(r.Players) == 2 {
				r.state.Objects = CreateObject(r)
				for _, player := range r.Players {
					player.score += 5
					tmp := r.state.Players[player.number]
					tmp.Score = player.score
					r.state.Players[player.number] = tmp
				}
			}

		case <-r.ticker.C:
			if len(r.Players) == 2 {

				tmp := r.state.Objects
				r.state.Objects = make([]ObjectState, 0)

				for index, object := range tmp {
					object.X = object.X - object.Speed
					object.ID = index

					if object.X > -30 {

						sumOfRad := 22.0 // sum of radiuses
						for _, player := range r.state.Players {
							centres := math.Sqrt(float64(((player.X - object.X) * (player.X - object.X)) + ((player.Y - object.Y) * (player.Y - object.Y))))

							if centres < sumOfRad {
								r.state.Objects = append(r.state.Objects, object)

								r.mu.Lock()
								insertScoreToDB(r.Players)
								r.mu.Unlock()

								for _, p := range r.Players {
									p.SendState(r.state)

									if p.ID == player.ID {
										p.SendMessage(&models.Logs{Head: "GAME ENDED", Content: "LOST"})
									} else {
										p.SendMessage(&models.Logs{Head: "GAME ENDED", Content: "WON"})
									}

									r.mu.Lock()
									delete(r.Players, p.number)
									r.mu.Unlock()
								}
								MainGame.mu.Lock()
								MainGame.RemoveRoom(r)
								MainGame.mu.Unlock()
								return
							}

						}
						r.state.Objects = append(r.state.Objects, object)
					}

				}

				for _, player := range r.Players {
					player.SendState(r.state)
				}
			}
		}
	}
}
