package buffers

import "log"

var IncomingMessages chan []byte
var Topics map[string] chan float64
var Subscribers map[string] []string
var Notify chan string

func Initialize() {
	IncomingMessages = make(chan []byte, 100)
	Topics = make(map[string] chan float64)
	Notify = make(chan string, 100)
	Subscribers = make(map[string] []string)
	log.Println("Initialized message buffers..")
}