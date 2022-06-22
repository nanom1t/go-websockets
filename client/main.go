package main

import (
	"fmt"
	"log"
	"github.com/gorilla/websocket"
	"os"
	"net/url"
)

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	urlAddress := url.URL{
		Scheme: "ws",
		Host: fmt.Sprintf("%s:%s", host, port),
		Path: "/ws",
	}

	// connect
	connection, _, err := websocket.DefaultDialer.Dial(urlAddress.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	done := make(chan struct{})

	go func() {
		for {
			_, message, err := connection.ReadMessage()
			if err != nil {
				log.Println("ERR:", err)
				return
			}

			log.Printf("MSG: %s", message)

			if string(message) == "done" {
				done <- struct{}{}
			}
		}
	}()

	<- done

	fmt.Println("Done!")
}
