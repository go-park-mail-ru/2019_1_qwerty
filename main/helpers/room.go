package helpers

import (
	"2019_1_qwerty/models"
	"math"
	"math/rand"
	"sync"
	"time"
)

//PlayerState - player state in room->game
type PlayerState struct {
	ID string
	X  int
	Y  int
}

//ObjectState - object state in room->game
type ObjectState struct {
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
func NewRoom(maxPlayers int) *Room {
	return &Room{
		MaxPlayers: maxPlayers,
		register:   make(chan *Player),
		unregister: make(chan *Player),
		ticker:     time.NewTicker(25 * time.Millisecond),   // HAS TO BEE 100 * ...
		workTicker: time.NewTicker(1500 * time.Millisecond), // meteors
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
	object := ObjectState{X: 350, Y: rand.Intn(120-30) + 30, Speed: rand.Intn(10-4) + 4}
	r.state.Objects = append(r.state.Objects, object)
	return r.state.Objects
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
			}

		case <-r.ticker.C:
			if len(r.Players) == 2 {

				tmp := r.state.Objects
				r.state.Objects = make([]ObjectState, 0)

				for _, object := range tmp {
					object.X = object.X - object.Speed

					if object.X > -30 {

						sumOfRad := 34.0 // sum of radiuses
						for _, player := range r.state.Players {
							centres := math.Sqrt(float64(((player.X - object.X) * (player.X - object.X)) + ((player.Y - object.Y) * (player.Y - object.Y))))

							if centres < sumOfRad {
								r.state.Objects = append(r.state.Objects, object)
								for _, p := range r.Players {
									p.SendState(r.state)
									r.mu.Lock()
									delete(r.Players, p.number)
									r.mu.Unlock()
									p.SendMessage(&models.Logs{Head: "GAME ENDED", Content: nil})
								}
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
