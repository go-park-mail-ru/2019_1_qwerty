package helpers

import (
	"sync"
)

//Game - main game struct
type Game struct {
	MaxRooms int
	mu       *sync.Mutex
	rooms    []*Room
	register chan *Player
}

//MainGame - main game
var MainGame = NewGame(10000)

func init() {
	go MainGame.Run()
}

//NewGame - creates new game
func NewGame(maxRooms int) *Game {
	return &Game{
		register: make(chan *Player),
		MaxRooms: maxRooms,
	}
}

//AddRoom - adds room to game
func (g *Game) AddRoom(room *Room) {
	g.rooms = append(g.rooms, room)
}

//AddPlayer - add player to room
func (g *Game) AddPlayer(player *Player) {
	g.register <- player
}

//Run - run game
func (g *Game) Run() {

LOOP:
	for {
		player := <-g.register
		for _, room := range g.rooms {
			if len(room.Players) < room.MaxPlayers {
				if len(room.Players) == 0 {
					player.number = "player1"
				} else {
					player.number = "player2"
				}
				go room.Run()
				room.mu.Lock()
				room.AddPlayer(player)
				room.mu.Unlock()
				continue LOOP
			}
		}

		room := NewRoom(2)
		room.mu.Lock()
		g.AddRoom(room)
		room.mu.Unlock()
		go room.Run()
		room.mu.Lock()
		player.number = "player1"
		room.AddPlayer(player)
		room.mu.Unlock()
	}
}
