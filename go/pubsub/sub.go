// sub.go
package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	// Create a new SUB socket
	subSocket, err := zmq.NewSocket(zmq.SUB)
	if err != nil {
		fmt.Println("Couldn't create SUB socket:", err)
		return
	}
	defer subSocket.Close()

	// Connect the SUB socket to the PUB socket
	err = subSocket.Connect("tcp://127.0.0.1:5680")
	if err != nil {
		fmt.Println("Couldn't connect SUB socket:", err)
		return
	}

	// Subscribe to messages with prefix "1"
	subSocket.SetSubscribe("1")

	// Start receiving messages
	for {
		message, err := subSocket.Recv(0)
		if err != nil {
			fmt.Println("Couldn't receive message:", err)
			break
		}
		fmt.Println("Received:", message)
	}
}
