package requestHandlers

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

func (crh ClientRequestHandler) SendAndReceive(msgToSend []byte) []byte {
	// stablish socket connection
	conn, err := net.Dial("tcp", crh.GetAddr())
	if (err != nil) {
		log.Fatal("Error establishing TCP connection with server. ", err)
	}

	// send message
	_, err = conn.Write(msgToSend)
	if (err != nil) {
		log.Fatal("Error sending message to server. ", err)
	}

	// wait for response and return
	responseMsg := make([]byte, 8)
	_, err = conn.Read(responseMsg)
	if (err != nil) {
		log.Fatal("Error receiving message from server. ", err)
	}

	return responseMsg
}