package infrastructure

import (
	"log"
	"net"
	"strconv"
)

type ClientRequestHandler struct {
	ServerHost string
	ServerPort int
}

func (crh ClientRequestHandler) GetAddr() string {
	return crh.ServerHost + ":" + strconv.Itoa(crh.ServerPort)
}

func (crh ClientRequestHandler) StablishConnection() net.Conn {
	conn, err := net.Dial("tcp", crh.GetAddr())
	if (err != nil) {
		log.Fatal("Error stablishing tcp connection ", err)
	}

	return conn
}

func (crh ClientRequestHandler) SendAndReceive(msgToSend []byte, conn net.Conn) []byte {
	// send message
	_, err := conn.Write(msgToSend)
	if (err != nil) {
		log.Fatal("Error sending message to server. ", err)
	}

	// wait for response and return
	responseMsg := make([]byte, 512)
	_, err = conn.Read(responseMsg)
	if (err != nil) {
		log.Fatal("Error receiving message from server. ", err)
	}

	return responseMsg
}