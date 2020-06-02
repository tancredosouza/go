package buffers

import "log"

var IncomingMessages chan []byte

func Initialize() {
	IncomingMessages = make(chan []byte, 100)
	log.Println("Initialized message buffers..")
}