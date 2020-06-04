package buffers

import (
	"../locks"
	"../persist"
	"log"
	"os"
)

var IncomingMessages chan []byte
var OutgoingMessages map[string] [][]byte
var Topics map[string] chan float64
var Subscribers map[string] []string
var ToNotifyTopicNames chan string

func Initialize() {
	IncomingMessages = make(chan []byte, 100)
	OutgoingMessages = make(map[string] [][]byte)
	ToNotifyTopicNames = make(chan string, 100)
	Topics = make(map[string] chan float64)
	Subscribers = make(map[string] []string)

	filepath := "./database/subscribers"
	if (fileExists(filepath)) {
		log.Println("Loading persisted subscriptions")

		locks.SubscribersLock.Lock()
		err := persist.Load(filepath, &Subscribers)
		if (err != nil) {
			log.Fatal("Error loading subscribers database", err)
		}
		locks.SubscribersLock.Unlock()
	}

	log.Println("Initialized message buffers..")
}

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}