package main

import (
	"../infrastructure"
	"log"
)

func main() {
	crh := infrastructure.ClientRequestHandler{
		ServerHost: "localhost",
		ServerPort: 3993,
	}

	msg := crh.SendAndReceive([]byte("Ola servidor"))
	log.Println(string(msg))
}
