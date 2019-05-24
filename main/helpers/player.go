package helpers

import (
	"2019_1_qwerty/models"
	"log"

	"github.com/gorilla/websocket"
)

//Player - player struct
type Player struct {
	conn *websocket.Conn
	ID   string
	room *Room
	in   chan *models.Logs
	out  chan *models.Logs
}

//NewPlayer - creates new player
func NewPlayer(connection *websocket.Conn, id string) *Player {
	return &Player{
		ID:   id,
		in:   make(chan *models.Logs),
		out:  make(chan *models.Logs),
		conn: connection,
	}
}

func handleCoordinates(player *Player, message *models.Logs) {
	tmp := player.room.state.Players[player.ID]
	switch message.Head {
	case "LEFT":
		tmp.X = tmp.X - 1
	case "RIGHT":
		tmp.X = tmp.X + 1
	case "UP":
		tmp.Y = tmp.Y - 1
	case "DOWN":
		tmp.Y = tmp.Y + 1
	}
	player.room.state.Players[player.ID] = tmp
}

//Listen - listen info
func (p *Player) Listen() {
	go func() {
		for {
			message := &models.Logs{}
			err := p.conn.ReadJSON(message)
			log.Println(err)
			if websocket.IsUnexpectedCloseError(err) {
				p.room.RemovePlayer(p)
				//for _, player := range p.room.Players {
				log.Println(p.ID, "disconnected from the server")
				//player.room.RemovePlayer(player)
				//}

				return
			}
			if err != nil {
				//log.Println("cant read json")
				continue
			}
			p.in <- message
		}
	}()

	for {
		select {
		case message := <-p.out:
			p.conn.WriteJSON(message)
		case message := <-p.in:
			handleCoordinates(p, message)
		}
	}
}

//SendState - send info to front about player
func (p *Player) SendState(state *RoomState) {
	p.out <- &models.Logs{Head: "STATE", Content: state.Players}
	p.out <- &models.Logs{Head: "OBJECTS", Content: state.Objects}
}

//SendMessage - send info to front about player
func (p *Player) SendMessage(message *models.Logs) {
	p.out <- message
}
