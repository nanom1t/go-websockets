package main

import (
	"fmt"
	"net/http"
	"time"
	"log"
	"github.com/gorilla/websocket"
	"os"
)

var upgrader = websocket.Upgrader{}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer connection.Close()

	for {
		messageType, message, err := connection.ReadMessage()
		if err != nil {
			log.Println("ERR:", err)
			break
		}

		log.Printf("MESSAGE: %s", message)

		err = connection.WriteMessage(messageType, message)
		if err != nil {
			log.Println("ERR:", err)
			break
		}
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	server := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", host, port),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ws", websocketHandler)

	log.Printf("Starting server on %s:%s", host, port)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
