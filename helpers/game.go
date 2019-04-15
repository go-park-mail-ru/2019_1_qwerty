package helpers

import (
	"fmt"
	"sync"
)

//Game - main game struct
type Game struct {
	MaxRooms uint
	mu       *sync.Mutex
	rooms    []*Room
	register chan *Player
}

//MainGame - main game
var MainGame = NewGame(10)

func init() {
	go MainGame.Run()
}

//NewGame - creates new game
func NewGame(maxRooms int) *Game {
	return &Game{
		register: make(chan *Player),
	}
}

//AddRoom - adds room to game
func (g *Game) AddRoom(room *Room) {
	g.rooms = append(g.rooms, room)
}

//AddPlayer - add player to room
func (g *Game) AddPlayer(player *Player) {
	fmt.Println(player.ID, "queued to add")
	g.register <- player
}

//Run - run game
func (g *Game) Run() {
	fmt.Println("main loop started")
LOOP:
	for {
		player := <-g.register

		for _, room := range g.rooms {
			if len(room.Players) < int(room.MaxPlayers) {
				room.AddPlayer(player)
				continue LOOP
			}
		}

		room := NewRoom(2)
		g.AddRoom(room)
		go room.Run()
		room.AddPlayer(player)
	}
}
