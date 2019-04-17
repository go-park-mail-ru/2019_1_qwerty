package helpers

import (
	"testing"
)

func TestNewGame(t *testing.T) {
	input := uint(5)
	game := NewGame(input)

	if game.MaxRooms != input {
		t.Fatalf("got %d, expected %d", game.MaxRooms, input)
	}
	game.AddRoom(&Room{})

}
