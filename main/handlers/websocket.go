package handlers

import (
	"2019_1_qwerty/helpers"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

//WebsocketConn - upgrade to websocket
func WebsocketConn(w http.ResponseWriter, r *http.Request) {
	upgrader := &websocket.Upgrader{}
	cookie, err := r.Cookie("sessionid")

	if err != nil {
		log.Println("no auth!")
		ErrorMux(&w, r, http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, http.Header{"Upgrade": []string{"websocket"}})

	if err != nil {
		ErrorMux(&w, r, http.StatusInternalServerError)
		log.Println("got error while connecting", err)
		return
	}

	log.Println("connected!")

	player := helpers.NewPlayer(conn, cookie.Value)
	go player.Listen()
	helpers.MainGame.AddPlayer(player)
}
