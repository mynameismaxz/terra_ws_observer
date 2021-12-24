package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/gorilla/websocket"
)

type subscribeData struct {
	Subscribe string `json:"subscribe"`
	ChainId   string `json:"chain_id"`
}

func main() {
	interrupt := make(chan os.Signal, 1)

	u := url.URL{Scheme: "wss", Host: "observer.terra.dev"}
	fmt.Printf("Connected to %s...\n", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		panic(err)
	}
	sub := &subscribeData{
		Subscribe: "new_txs",
		ChainId:   "columbus-5",
	}

	payload, err := json.Marshal(sub)
	if err != nil {
		panic(err)
	}
	msg := string(payload)
	err = c.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		panic(err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				panic(err)
				return
			}
			fmt.Printf("recv: %s\n\n", msg)
		}
	}()

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			fmt.Println("Interrupt mode.")
			return
		}
	}
}
