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

	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("Message #%d", i)

		err = connection.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("ERR:", err)
			break
		}

		time.Sleep(time.Second)
	}

	err = connection.WriteMessage(websocket.TextMessage, []byte("done"))
	if err != nil {
		log.Println("ERR:", err)
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
