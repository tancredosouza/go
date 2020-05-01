package main

import (
	"../queue"
	"../stack"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
)

func main() {
	queue := new(queue.QueueRPC)
	stack := new(stack.StackRPC)

	server := rpc.NewServer()
	err := server.RegisterName("app.Queue", queue)
	if (err != nil) {
		log.Fatal("Error occurred when registering name: ", err)
	}

	err = server.RegisterName("app.Stack", stack)
	if (err != nil) {
		log.Fatal("Error occurred when registering name: ", err)
	}

	server.RegisterName("app.Stack", stack)

	go func() {
		l, err := net.Listen("tcp", "localhost:1567")
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		server.Accept(l)
	}()

	fmt.Scanln()
}