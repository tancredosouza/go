package requestHandlers

import (
	"log"
	"net"
	"strconv"
)

type ServerRequestHandler struct {
	ServerHost string
	ServerPort int
}

var listener net.Listener
var err error
var clientConn net.Conn

func (srh ServerRequestHandler) StartListening() {
	listener, err = net.Listen("tcp", srh.GetAddr())
	if (err != nil) {
		log.Fatal("Error while creating listener. ", err)
	}
}

func (srh ServerRequestHandler) GetAddr() string {
	return srh.ServerHost + ":" + strconv.Itoa(srh.ServerPort)
}

func (srh ServerRequestHandler) Receive() []byte {
	// get message
	clientConn, err = listener.Accept()
	if (err != nil) {
		log.Fatal("Error while accepting client connection. ", err)
	}

	// return message
	clientMsg := make([]byte, 8)
	_, err = clientConn.Read(clientMsg)
	if (err != nil) {
		log.Fatal("Error while reading message from client. ", err)
	}

	return clientMsg
}

func (srh ServerRequestHandler) Send(msg []byte) {
	// send message
	_, err = clientConn.Write(msg)
	if (err != nil) {
		log.Fatal("Error writing response to client. ", err)
	}

	clientConn.Close()
}