package infrastructure

import (
	"log"
	"net"
	"strconv"
)

type ServerRequestHandler struct {
	ServerHost     string
	ServerPort     int
	Conn           net.Conn
	listener       net.Listener
}

var receivedBuffer chan []byte
var toSendBuffer   chan []byte

func (srh *ServerRequestHandler) Initialize() {
	receivedBuffer = make(chan []byte, 10)
	toSendBuffer = make(chan []byte, 10)

	srh.StartListening()
}

func (srh *ServerRequestHandler) Send(msgToSend []byte) {
	toSendBuffer <- msgToSend
}

func (srh *ServerRequestHandler) Receive() []byte {
	msgReceived := <- receivedBuffer
	return msgReceived
}

func (srh *ServerRequestHandler) keepSending() {
	for {
		msgToSend := <- toSendBuffer
		err := Send(msgToSend, srh.Conn)

		if (err != nil) {
			log.Println("error while sending -> ", string(msgToSend), err)
			break
		}
	}
}

func (srh *ServerRequestHandler) keepReceiving() {
	for {
		receivedData, err := Receive(srh.Conn)

		if (err != nil) {
			log.Println("error while receiving -> ", err)
			break
		}

		if (receivedData == nil) {
			continue
		}

		receivedBuffer <- receivedData
	}
}

/*
As-is, for each connection the code creates a listener and
then closes it when the data is sent. This could harm execution
time. To solve this problem, the function below could be
used instead. However, this would require the Distribution
Layer to explicitly start the SRH, breaking the management
isolation provided by the layered architecture.
*/
func (srh *ServerRequestHandler) StartListening() {
	srh.listener, _ = net.Listen("tcp", srh.GetAddr())
	log.Println("Listening...")
}

func (srh *ServerRequestHandler) GetAddr() string {
	return srh.ServerHost + ":" + strconv.Itoa(srh.ServerPort)
}

func (srh *ServerRequestHandler) AcceptNewConnection() {
	var err error
	srh.Conn, err = srh.listener.Accept()

	if (err != nil) {
		log.Fatal("Error while accepting connection ", err)
	}

	log.Println("Successfully accepted connection")
}

func (srh *ServerRequestHandler) CreateQueues() {
	go srh.keepSending()
	go srh.keepReceiving()
}

func (srh *ServerRequestHandler) StopListening() {
	err := srh.listener.Close()

	if (err != nil) {
		log.Fatal("Error closing listener. ", err)
	}
	log.Println("Closed listener")
}