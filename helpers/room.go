package helpers

import (
	"2019_1_qwerty/models"
	"fmt"
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
	ID string
	X  int
	Y  int
}

//RoomState - room model state
type RoomState struct {
	Players     []PlayerState
	Objects     []ObjectState
	CurrentTime time.Time
}

//Room - room model
type Room struct {
	ID         string
	MaxPlayers uint
	Players    map[string]*Player
	mu         *sync.Mutex
	register   chan *Player
	unregister chan *Player
	ticker     *time.Ticker
	state      *RoomState
}

//NewRoom - create new room
func NewRoom(maxPlayers uint) *Room {
	return &Room{
		MaxPlayers: maxPlayers,
		register:   make(chan *Player),
		unregister: make(chan *Player),
		ticker:     time.NewTicker(1 * time.Second),
		Players:    make(map[string]*Player),
		mu:         new(sync.Mutex),
	}
}

//AddPlayer - add player to slice of players
func (r *Room) AddPlayer(player *Player) {
	player.room = r
	r.register <- player
}

//RemovePlayer - add player to slice of players
func (r *Room) RemovePlayer(player *Player) {
	r.unregister <- player
}

//Run - runs rooms
func (r *Room) Run() {
	fmt.Println("room loop started")
	for {
		select {
		case player := <-r.unregister:
			r.mu.Lock()
			delete(r.Players, player.ID)
			r.mu.Unlock()
			fmt.Println(player.ID, "was deleted from the room")
		case player := <-r.register:
			r.mu.Lock()
			r.Players[player.ID] = player
			r.mu.Unlock()
			fmt.Println(player.ID, "joined the room")
			player.SendMessage(&models.Logs{Head: "CONNECTED", Content: nil})
		case <-r.ticker.C:
			for _, player := range r.Players {
				player.SendState(r.state)
			}
		}
	}
}
