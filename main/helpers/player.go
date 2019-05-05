package helpers

import (
	"2019_1_qwerty/models"
	"fmt"

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

//Listen - listen info
func (p *Player) Listen() {
	go func() {
		for {
			message := &models.Logs{}
			err := p.conn.ReadJSON(message)
			if websocket.IsUnexpectedCloseError(err) {
				p.room.RemovePlayer(p)
				fmt.Println(p.ID, "disconnected from the server")
				return
			}
			if err != nil {
				fmt.Println("cant read json")
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
			fmt.Println("income: ", message)

		}
	}
}

//SendState - send info to front about player
func (p *Player) SendState(state *RoomState) {
	p.out <- &models.Logs{Head: "STATE", Content: state}
}

//SendMessage - send info to front about player
func (p *Player) SendMessage(message *models.Logs) {
	p.out <- message
}
