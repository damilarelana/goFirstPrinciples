package main

import (
	"fmt"
	"time"
)

// The Pinger function, receives the `msgChannel` channel as an argument
func pinger(msgChannel chan<- string) {
	for i := 0; ; i++ { // an infinite ever increasing counter :
		msgChannel <- "ping" // sends a "ping" string to the channel
	}
}

// The Ponger function, receives the `msgChannel` channel as an argument
func ponger(msgChannel chan<- string) {
	for i := 0; ; i++ { // an infinite ever increasing counter :
		msgChannel <- "pong" // sends a "pong" string to the channel
	}
}

// The Printer, receives the `msgChannel` channel as an argument
func printer(msgChannel <-chan string) {
	for { // this is always true, hence is similar to "always listening"
		message := <-msgChannel // recieves a string from the channel
		fmt.Println(message)
		time.Sleep(time.Second * 1)
	}
}

func main() {

	msgChannel := make(chan string, 5) // declare/initialize a bufferred channel

	// create the goroutines from initial functions
	go pinger(msgChannel)
	go ponger(msgChannel)
	go printer(msgChannel)

	// create an artificial pause "so as to ensure the main functoin does not exit"
	var tempString string
	fmt.Scanln(&tempString)
}
