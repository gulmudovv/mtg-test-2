package main

import (
	"MTG-test-2/server/database"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsHandle(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	// helpful log statement to show connections
	log.Println("Client Connected", ws.RemoteAddr())

	go reader(ws)
}

func reader(conn *websocket.Conn) {
	for {
		// read in a message
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		//save to database
		id := conn.RemoteAddr().String()
		data, err := decode(string(p))
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("Data %s id %s", data, id)
		database.Create(id, data)

	}
}

func main() {
	database.ConnectDB()
	http.HandleFunc("/ws", wsHandle)
	log.Println("http server started on :8585")
	err := http.ListenAndServe(":8585", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func decode(enc string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(enc)
	if err != nil {
		fmt.Println("decode error:", err)
		return "", err
	}
	return string(decoded), nil
}
