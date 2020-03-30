package main

// Connect, subscribe on channel, publish into channel, read presence and history info.

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/centrifugal/centrifuge-go"
)

type testMessage struct {
	Input string `json:"input"`
}

func main() {

	url := "ws://localhost:8000/connection/websocket"

	c := centrifuge.New(url, centrifuge.DefaultConfig())
	defer c.Close()

	err := c.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	sub, err := c.NewSubscription("global")
	if err != nil {
		log.Fatalln(err)
	}

	err = sub.Subscribe()
	if err != nil {
		log.Fatalln(err)
	}

	history, err := sub.History()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%d messages in channel %s history", len(history), sub.Channel())

	presence, err := sub.Presence()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%d clients in channel %s", len(presence), sub.Channel())

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter message: ")
		text, _ := reader.ReadString('\n')
		data := testMessage{Input: text}
		dataBytes, _ := json.Marshal(data)
		err = sub.Publish(dataBytes)
		if err != nil {
			fmt.Println("failed to publish message")
		} else {
			fmt.Println("Successfully publish a message")
		}
	}
}
