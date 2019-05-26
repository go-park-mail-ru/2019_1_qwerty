package handlers

import (
	"2019_1_qwerty/helpers"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

//WebsocketConn - upgrade to websocket
func WebsocketConn(w http.ResponseWriter, r *http.Request) {
	val, _ := r.URL.Query()["nickname"]
	cookie, _ := helpers.GetID(val[0])

	if cookie == "" {
		log.Println("no auth!")
		ErrorMux(&w, r, http.StatusUnauthorized)
		return
	}

	upgrader := &websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, http.Header{"Upgrade": []string{"websocket"}})

	if err != nil {
		ErrorMux(&w, r, http.StatusInternalServerError)
		log.Println("got error while connecting", err)
		return
	}

	log.Println("connected!")

	player := helpers.NewPlayer(conn, cookie)
	go player.Listen()
	helpers.MainGame.AddPlayer(player)
}
