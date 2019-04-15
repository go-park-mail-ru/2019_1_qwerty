package handlers

import (
	"2019_1_qwerty/helpers"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

//WebsocketConn - upgrade to websocket
func WebsocketConn(w http.ResponseWriter, r *http.Request) {
	upgrader := &websocket.Upgrader{}
	cookie, err := r.Cookie("sessionid")

	if err != nil {
		fmt.Println("no auth!")
		w.WriteHeader(401)
		return
	}

	conn, err := upgrader.Upgrade(w, r, http.Header{"Upgrade": []string{"websocket"}})

	if err != nil {
		fmt.Println("got error while connecting", err)
		w.WriteHeader(500)
		return
	}

	fmt.Println("connected!")

	player := helpers.NewPlayer(conn, cookie.Value)
	go player.Listen()
	helpers.MainGame.AddPlayer(player)
}
