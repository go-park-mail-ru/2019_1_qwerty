package helpers

import (
	"testing"

	"github.com/gorilla/websocket"
)

func TestNewPlayer(t *testing.T) {
	var connection *websocket.Conn
	input := "player"
	player := NewPlayer(connection, input)

	if player.ID != input {
		t.Fatalf("got %s, expected %s", player.ID, input)
	}
	// player.SendMessage(&models.Logs{})
	// player.Listen()
}
