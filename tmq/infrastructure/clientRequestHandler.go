package infrastructure

import (
	"log"
	"net"
	"strconv"
)

type ClientRequestHandler struct {
	ServerHost string
	ServerPort int
	conn       net.Conn
}

var err error

func (crh *ClientRequestHandler) GetAddr() string {
	return crh.ServerHost + ":" + strconv.Itoa(crh.ServerPort)
}

func (crh *ClientRequestHandler) Initialize() {
	log.Println("Initializing client connection")
	for {
		crh.conn, err = net.Dial("tcp", crh.GetAddr())
		if err == nil {
			break
		}
		log.Println(err)
	}
}

func (crh *ClientRequestHandler) Send(msgToSend []byte) {
	err = Send(msgToSend, crh.conn)
	if (err != nil) {
		log.Fatal(err)
	}
}

/*
func (crh *ClientRequestHandler) SendAndReceive(msgToSend []byte) []byte {
	crh.Initialize()

	err = Send(msgToSend, crh.conn)
	if (err != nil) {
		log.Fatal(err)
	}

	responseMsg, err := Receive(crh.conn)
	if (err != nil) {
		log.Fatal(err)
	}

	return responseMsg
}
*/

func (crh *ClientRequestHandler) CloseConnection() {
	crh.conn.Close()
}