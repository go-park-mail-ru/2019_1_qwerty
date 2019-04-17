package helpers

import "testing"

func TestNewRoom(t *testing.T) {
	input := uint(2)
	room := NewRoom(input)

	if room.MaxPlayers != input {
		t.Fatalf("got %d, expected %d", room.MaxPlayers, input)
	}

}
