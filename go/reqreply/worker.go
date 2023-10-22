package main

import (
	"encoding/json"
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

type Message struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	zctx, _ := zmq.NewContext()

	s, _ := zctx.NewSocket(zmq.SUB)
	s.Connect("tcp://localhost:5757")
	s.SetSubscribe("dev.to")

	for {
		address, err := s.Recv(0)
		if err != nil {
			panic(err)
		}

		if msg, err := s.Recv(0); err != nil {
			panic(err)
		} else {
			contents := Message{}
			if err := json.Unmarshal([]byte(msg), &contents); err != nil {
				panic(err)
			}
			fmt.Println("Received message from " + address + " channel.")
			fmt.Printf("%+v\n", contents)
		}
	}
}
