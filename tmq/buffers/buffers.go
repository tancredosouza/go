package buffers

import (
	"../persist"
	"log"
	"os"
)

var IncomingMessages chan []byte
var Topics map[string] chan float64
var Subscribers map[string] []string
var ToNotifyTopicNames chan string

func Initialize() {
	IncomingMessages = make(chan []byte, 100)
	ToNotifyTopicNames = make(chan string, 100)
	Topics = make(map[string] chan float64)
	Subscribers = make(map[string] []string)

	filepath := "./database/subscribers"
	if (fileExists(filepath)) {
		log.Println("Loading persisted subscriptions")

		err := persist.Load(filepath, &Subscribers)
		if (err != nil) {
			log.Fatal("Error loading subscribers database", err)
		}
	}

	log.Println("Initialized message buffers..")
}

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// path/to/whatever does not exist
		return false
	}
	return true
}