package helpers

import (
	"2019_1_qwerty/models"

	"github.com/gorilla/websocket"
)

//Player - player struct
type Player struct {
	conn   *websocket.Conn
	ID     string
	number string
	room   *Room
	in     chan *models.Logs
	out    chan []byte
	score  int
}

//NewPlayer - creates new player
func NewPlayer(connection *websocket.Conn, key string, value string) *Player {
	return &Player{
		ID:     key,
		number: value,
		in:     make(chan *models.Logs),
		out:    make(chan []byte),
		conn:   connection,
		score:  0,
	}
}

func handleCoordinates(player *Player, message *models.Logs) {
	tmp := player.room.state.Players[player.number]
	switch message.Head {
	case "LEFT":
		tmp.X = tmp.X - 2
	case "RIGHT":
		tmp.X = tmp.X + 2
	case "UP":
		tmp.Y = tmp.Y - 2
	case "DOWN":
		tmp.Y = tmp.Y + 2
	}
	player.room.state.Players[player.number] = tmp
}

//Listen - listen info
func (p *Player) Listen() {
	go func() {
		for {
			message := &models.Logs{}
			err := p.conn.ReadJSON(message)

			if websocket.IsUnexpectedCloseError(err) {
				for _, player := range p.room.Players {
					player.room.RemovePlayer(player)
				}
				return
			}

			if err != nil {
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
	jsonState := &models.Logs{Head: "STATE", Content: state.Players}
	jsonObjects := &models.Logs{Head: "OBJECTS", Content: state.Objects}
	resultState, _ := jsonState.MarshalJSON()
	resultObjects, _ := jsonObjects.MarshalJSON()
	p.out <- resultState
	p.out <- resultObjects
}

//SendMessage - send info to front about player
func (p *Player) SendMessage(message *models.Logs) {
	result, _ := message.MarshalJSON()
	p.out <- result
}
