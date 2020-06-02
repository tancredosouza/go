package buffers

import "log"

var IncomingMessages chan []byte
var Topics map[string] chan []byte

func Initialize() {
	IncomingMessages = make(chan []byte, 100)
	Topics = make(map[string] chan []byte)
	log.Println("Initialized message buffers..")
}