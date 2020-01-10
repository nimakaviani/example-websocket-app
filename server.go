package main

// +build ignore

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"os"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		log.Printf("recv: %s", message)

		host, err := os.Hostname()
		if err != nil {
			log.Println("read: ", err)
		}

		pongMsg := fmt.Sprintf("%s: pong", time.Now())
		err = c.WriteMessage(mt, []byte(pongMsg))
		if err != nil {
			log.Println("write:", err)
			break
		}
		err = c.WriteMessage(mt, []byte(fmt.Sprintf("host: %s", host)))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", echo)

	log.Println("starting ....")
	log.Fatal(http.ListenAndServe(*addr, nil))
}
