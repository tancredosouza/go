package infrastructure

import (
	"log"
	"net"
	"strconv"
)

type ClientRequestHandler struct {
	ServerHost string
	ServerPort int
	receivedBuffer chan []byte
	toSendBuffer chan []byte
	conn net.Conn
}

func (crh ClientRequestHandler) GetAddr() string {
	return crh.ServerHost + ":" + strconv.Itoa(crh.ServerPort)
}

func (crh *ClientRequestHandler) Initialize() {
	crh.receivedBuffer = make(chan []byte, 10)
	crh.toSendBuffer = make(chan []byte, 10)

	log.Println("Trying to dial ", crh.GetAddr())
	var err error
	crh.conn, err = net.Dial("tcp", crh.GetAddr())
	if (err != nil) {
		log.Fatal("Error establishing tcp connection ", err)
	}

	log.Println("Connection established! Hooray.")

	go crh.keepSending()
	go crh.keepReceiving()
}

func (crh *ClientRequestHandler) Send(msgToSend []byte) {
	crh.toSendBuffer <- msgToSend
}

func (crh *ClientRequestHandler) Receive() []byte {
	return <- crh.receivedBuffer
}

func (crh *ClientRequestHandler) keepSending() {
	for {
		msgToSend := <- crh.toSendBuffer
		err := Send(msgToSend, crh.conn)

		if (err != nil) {
			log.Println("error while sending -> ", err)
			break
		}
	}
}

func Send(msgToSend []byte, conn net.Conn) error {
	// send message
	_, err := conn.Write(msgToSend)
	return err
}

func (crh *ClientRequestHandler) keepReceiving() {
	for {
		receivedData, err := Receive(crh.conn)

		if (err != nil) {
			log.Println("error while receiving -> ", err)
			break
		}

		crh.receivedBuffer <- receivedData
	}
}

func Receive(conn net.Conn) ([]byte, error) {
	// wait for response and return
	responseMsg := make([]byte, 512)
	_, err := conn.Read(responseMsg)

	return responseMsg, err
}

func (crh *ClientRequestHandler) CloseConnection() {
	crh.conn.Close()
}