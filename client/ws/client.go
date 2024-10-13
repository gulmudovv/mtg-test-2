package ws

import (
	"encoding/base64"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var done chan struct{}
var interrupt chan os.Signal

func receiveHandler(connection *websocket.Conn) {
	defer close(done)
	for {
		_, msg, err := connection.ReadMessage()
		if err != nil {
			log.Println("Error in receive:", err)
			return
		}
		log.Printf("Received: %s\n", msg)
	}
}

func Worker(num int) {

	interrupt = make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	done = make(chan struct{})

	u := url.URL{Scheme: "ws", Host: "localhost:8585", Path: "/ws"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Error connecting to Websocket Server:", err)
	}
	defer conn.Close()

	go receiveHandler(conn)

	// We send our relevant packets here
	for {
		select {
		case <-time.After(time.Duration(1) * time.Millisecond * 1000):
			// Send an echo packet every second
			err := conn.WriteMessage(websocket.TextMessage, []byte("Hello from GolangDocs!"))
			if err != nil {
				log.Println("Error during writing to websocket:", err)
				return
			}

		case <-interrupt:
			// We received a SIGINT (Ctrl + C). Terminate gracefully...
			log.Println("Received SIGINT interrupt signal. Closing all pending connections")

			// Close our websocket connection
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Error during closing websocket:", err)
				return
			}

			select {
			case <-done:
				log.Println("Receiver Channel Closed! Exiting....")
			case <-time.After(time.Duration(1) * time.Second):
				log.Println("Timeout in closing receiving channel. Exiting....")
			}
			return
		default:
			message := answer()
			conn.WriteMessage(websocket.TextMessage, []byte(message))
		}
	}

}

func answer() string {
	// count := rand.Intn(10005-1000) + 1000
	// arr := make([]byte, count)

	// for i := 0; i < len(arr); i++ {
	// 	arr[i] = byte(rand.Intn(255))
	// }

	msg := base64.StdEncoding.EncodeToString([]byte("1234567890"))

	return msg
}
