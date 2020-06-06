package buffers

import (
	"log"
	"os"
	"../persist"
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

	loadPersistedSubscribers()
	loadPersistedOutgoingMessages()

	log.Println("Initialized message buffers..")
}

func loadPersistedSubscribers() {
	filepath := "./database/subscribers"
	if (fileExists(filepath)) {
		log.Println("Loading persisted subscriptions")

		err := persist.Load(filepath, &Subscribers)
		if (err != nil) {
			log.Fatal("Error loading subscribers database", err)
		}
	}
}

func loadPersistedOutgoingMessages() {
	filepath := "./database/outgoing"
	if (fileExists(filepath)) {
		log.Println("Loading persisted outgoing messages")

		err := persist.Load(filepath, &OutgoingMessages)
		if (err != nil) {
			log.Fatal("Error loading outgoing messages database", err)
		}
	}
}

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}