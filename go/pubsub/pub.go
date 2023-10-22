// pub.go
package main

import (
	"fmt"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	// Create a new PUB socket
	pubSocket, err := zmq.NewSocket(zmq.PUB)
	if err != nil {
		fmt.Println("Couldn't create PUB socket:", err)
		return
	}
	defer pubSocket.Close()

	// Bind the PUB socket to a TCP port
	err = pubSocket.Bind("tcp://127.0.0.1:5680")
	if err != nil {
		fmt.Println("Couldn't bind PUB socket:", err)
		return
	}

	// Start publishing messages
	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("1-Message %d", i)
		pubSocket.Send(message, 0)
		time.Sleep(time.Second)
	}
}
