package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/websocket"
)

func TestWebsocketConn(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(WebsocketConn))
	defer s.Close()

	d := websocket.Dialer{}
	_, _, err := d.Dial("ws://localhost:8080/api/ws", http.Header{"Cookie": []string{"sessionid=test"}})

	if err != nil {
		t.Fatalf("%v", err)
	}

	return

}
