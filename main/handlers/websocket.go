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
		Hits.WithLabelValues(string(401), r.URL.String()).Inc()
		FooCount.Add(1)
		w.WriteHeader(401)
		return
	}

	conn, err := upgrader.Upgrade(w, r, http.Header{"Upgrade": []string{"websocket"}})

	if err != nil {
		Hits.WithLabelValues(string(500), r.URL.String()).Inc()
		FooCount.Add(1)
		log.Println("got error while connecting", err)
		w.WriteHeader(500)
		return
	}

	log.Println("connected!")

	player := helpers.NewPlayer(conn, cookie.Value)
	go player.Listen()
	helpers.MainGame.AddPlayer(player)
}
