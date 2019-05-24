package helpers

import (
	"log"
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
	log.Println(player.ID, "queued to add")
	g.register <- player
}

//Run - run game
func (g *Game) Run() {

LOOP:
	for {
		player := <-g.register
		log.Println("game.go - new player in loop: ", player.ID)
		for _, room := range g.rooms {
			log.Println("got into rooms")
			if len(room.Players) < room.MaxPlayers {
				log.Println("len(room.Players) < room.MaxPlayers")
				room.AddPlayer(player)
				continue LOOP
			}
		}

		room := NewRoom(2)
		g.AddRoom(room)
		log.Println("room added: ", room)
		go room.Run()
		room.AddPlayer(player)
		log.Println("player added: ", player.ID)
	}
}
