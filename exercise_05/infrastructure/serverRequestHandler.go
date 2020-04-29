package infrastructure

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

/*
As-is, for each connection the code creates a listener and
then closes it when the data is sent. This could harm execution
time. To solve this problem, the function below could be
used instead. However, this would require the Distribution
Layer to explicitly start the SRH, breaking the management
isolation provided by the layered architecture.

func (srh ServerRequestHandler) StartListening() {
	listener, err = net.Listen("tcp", srh.GetAddr())
	if (err != nil) {
		log.Fatal("Error while creating listener. ", err)
	}
}
*/

func (srh ServerRequestHandler) GetAddr() string {
	return srh.ServerHost + ":" + strconv.Itoa(srh.ServerPort)
}

func (srh ServerRequestHandler) Receive() []byte {
	// get message
	listener, err = net.Listen("tcp", srh.GetAddr())
	if (err != nil) {
		log.Fatal("Error while creating listener. ", err)
	}

	clientConn, err = listener.Accept()
	if (err != nil) {
		log.Fatal("Error while accepting client connection. ", err)
	}

	// return message
	clientMsg := make([]byte, 1024)
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
	listener.Close()
}