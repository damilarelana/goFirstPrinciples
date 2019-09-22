package main

import (
	"fmt"
	"time"
)

// create the user defined Sleep function that recieves a time duration as argument
func bespokeSleepFunction(d time.Duration) {
	time.After(d)
}

func main() {

	// Create the message channels
	msgChannelOne := make(chan string, 5) // declare/initialize bufferred channel One
	msgChannelTwo := make(chan string, 5) // declare/initialize a bufferred channel Two

	// create the goroutines functions using anonymous functions
	go func() {
		for {
			msgChannelOne <- "from Channel 1"
			bespokeSleepFunction(time.Second * 2) // pause for 2 seconds after sending the initial message
		}
	}() // the () means we are calling the function immediately

	go func() {
		for {
			msgChannelTwo <- "from Channel 3"
			bespokeSleepFunction(time.Second * 3) // pause for 3 seconds after sending the initial message
		}
	}() // the () means we are calling the function immediately

	go func() {
		for {
			select {
			case messageOne := <-msgChannelOne:
				fmt.Println(messageOne)
			case messageTwo := <-msgChannelTwo:
				fmt.Println(messageTwo)
			case <-time.After(time.Second * 5):
				fmt.Println("timeout ... no message was sent")
			default:
				fmt.Println("nothing is ready or happening yet")
			}
		}
	}() // the () means we are calling the function immediately

	// create an artificial pause "so as to ensure the main functoin does not exit during goroutine runtime"
	var tempString string
	fmt.Scanln(&tempString)
}
