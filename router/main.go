package main

import (
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.OpenFile("../app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)

	for {
		logger.Println("This is a log message")
		time.Sleep(5 * time.Second) // Log every 5 seconds
	}
}
