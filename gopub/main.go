package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to the NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Define the message
	message := "This is a test message"

	// Publish the message to the "logs" topic
	err = nc.Publish("logs", []byte(message))
	if err != nil {
		log.Fatal(err)
	}

	// Ensure the message is sent before closing the connection
	nc.Flush()

	// Check if there are any errors during publishing
	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Published message to 'logs' topic: %s", message)

	// Wait for a short duration to ensure the message is received
	time.Sleep(1 * time.Second)
}
