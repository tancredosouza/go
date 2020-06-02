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

var conn net.Conn = nil
var err error

func (crh *ClientRequestHandler) GetAddr() string {
	return crh.ServerHost + ":" + strconv.Itoa(crh.ServerPort)
}

func (crh *ClientRequestHandler) Initialize() {
	log.Println("Initializing client connection")
	for {
		conn, err = net.Dial("tcp", crh.GetAddr())
		if err == nil {
			break
		}
		log.Println(err)
	}
}

func (crh *ClientRequestHandler) SendAndReceive(msgToSend []byte) []byte {
	crh.Initialize()

	err = Send(msgToSend, conn)
	if (err != nil) {
		log.Fatal(err)
	}

	responseMsg, err := Receive(conn)
	if (err != nil) {
		log.Fatal(err)
	}

	return responseMsg
}

func (ClientRequestHandler) CloseConnection() {
	conn.Close()
	conn = nil
}