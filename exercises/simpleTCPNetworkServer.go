package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

// Define the Server [meant do 3 things]
// 1. list and allow connections
// 2. respond to dynamic requests
// 3. server back content as the response
func mainServer() {
	// initiate protocols to listen for connection requests on a particular port
	initListen, err := net.Listen("tcp", ":9990")
	if err != nil {
		fmt.Println(err) // print the error message
		return
	}
	// proceed to accept connections after successfully listening
	for {
		initAccept, err := initListen.Accept() // accept the connection i.e. using the same listening interface we just initialized
		if err != nil {
			fmt.Println(err) // print the error message
			return
		}
		go handleServerConnection(initAccept) // proceed to now use the accepted connection "initAccept", by deploying it as goroutine
	}
}

// Define the Client [meant do 4 things]
// 1. initiate server connection
// 2. generate and encode messageContent
// 3. send encoded messageContent
// 4. close the established connection
func mainClient() {
	// initiated connections to server
	initConnect, err := net.Dial("tcp", "127.0.0.1:9990") // 127.0.0.1 is used here to avoid scenarios where localhost is redefined
	if err != nil {
		fmt.Println(err)
		return
	}
	// proceed to use the initiated connection to communicate e.g. send an encoded message
	messageContent := "This is a sample string, meant to be encoded and sent to Server"
	fmt.Println("Sending:", messageContent)                  // communicate the initiation of message encoding
	err = gob.NewEncoder(initConnect).Encode(messageContent) // use existing initConnect, encode the messageContent, do nothing with err :)
	initConnect.Close()                                      // close connection after sending encoded message
}

// Define the handshake-handler between the Client and Server
func handleServerConnection(newConnection net.Conn) {
	var messageContent string                                    // receive mesage from Client
	err := gob.NewDecoder(newConnection).Decode(&messageContent) // decode received message by referencing the pointer i.e. without passing it
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received:", messageContent) // communicate that message was received
	}
}

func main() {
	go mainServer() // run the server as a goroutine
	go mainClient() // run the client as a goroutine

	// create code termination hack using an expected input  that is not used for anything
	var unusedInput string
	fmt.Scanln(&unusedInput)
}
