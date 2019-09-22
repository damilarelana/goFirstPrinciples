package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// Server struct defined without any field set
type Server struct{}

// Inverse method mapped to Server struct
func (s *Server) Inverse(inputValue float64, returnedValue *float64) error {
	*returnedValue = 1 / inputValue
	return nil
}

// server where the rpc server lives/runs when called i.e. rpc server basically runs on a net server connection: extended by rpc package
func server() {
	rpc.Register(new(Server))                              // register Server instance, prior to accessing its method
	initServerListening, err := net.Listen("tcp", ":9990") // initialize the server's listening functionality
	if err != nil {
		fmt.Println(err) // print the error that caused the server not to start
		return
	}
	for { // a perpetual forloop to always accept incoming connections
		initAcceptConnection, err := initServerListening.Accept() //initialize acceptance of connetios
		if err != nil {
			continue // return to start of forloop i.e. continue accepting connections
		}
		go rpc.ServeConn(initAcceptConnection) // with a connection now established, this goroutine starts/runs rpc server (over the server connection) until the client hangs up
	}

}

// defining the rpcclient meant to interact with the rpc server
func client() {
	initRPCClient, err := rpc.Dial("tcp", "127.0.0.1:9990") //initialize rpc client
	if err != nil {
		fmt.Println(err)
		return
	}
	var returnedRPCValue float64                                                // initialize the variable we got back from the rpc
	err = initRPCClient.Call("Server.Inverse", float64(999), &returnedRPCValue) // finally making the method call and getting the returnedValue back at the address
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("via the RPC, the inverse of 999 is: ", returnedRPCValue)
	}
}

func main() {
	// now run the server and client concurrently
	go server()
	go client()

	// create code termination hack using an expected input  that is not used for anything
	var unusedInput string
	fmt.Scanln(&unusedInput)
}
