package main

import (
	"../requestHandlers"
	"log"
)

func main() {
	crh := requestHandlers.ClientRequestHandler{"localhost", 6966 }

	s := crh.SendAndReceive([]byte("olar"))

	log.Println(string(s))
}
