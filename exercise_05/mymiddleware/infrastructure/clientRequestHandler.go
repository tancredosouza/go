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

func (crh ClientRequestHandler) GetAddr() string {
	return crh.ServerHost + ":" + strconv.Itoa(crh.ServerPort)
}

func (crh ClientRequestHandler) SendAndReceive(msgToSend []byte) []byte {
	// stablish socket connection
	// connect to server
	if (conn == nil) {
		//log.Println("Initializing client connection")
		for {
			conn, err = net.Dial("tcp", crh.GetAddr())
			if err == nil {
				break
			}
			log.Println(err)
		}
	}
	//defer conn.Close()

	// send message
	_, err = conn.Write(msgToSend)
	if (err != nil) {
		log.Fatal("Error sending message to server. ", err)
	}

	// wait for response and return
	responseMsg := make([]byte, 512)
	_, err := conn.Read(responseMsg)
	if (err != nil) {
		log.Fatal("Error receiving message from server. ", err)
	}

	return responseMsg
}

func (ClientRequestHandler) CloseConnection() {
	conn.Close()
	conn = nil
}