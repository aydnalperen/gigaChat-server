package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("Error during connection upgrade: ", err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()

		if err != nil {
			log.Println("Error during message reading", err)
			break
		}
		log.Println("Received: ", message)

		err = conn.WriteMessage(messageType, message)

		if err != nil {
			log.Println("Error during message writing: ", err)
			break
		}
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index Page")
}

func main() {
	http.HandleFunc("/socket", SocketHandler)
	http.HandleFunc("/", Home)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
