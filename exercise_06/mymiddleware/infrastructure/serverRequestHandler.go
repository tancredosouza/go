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
var lock chan struct{}
var clientConnection net.Conn

/*
As-is, for each connection the code creates a listener and
then closes it when the data is sent. This could harm execution
time. To solve this problem, the function below could be
used instead. However, this would require the Distribution
Layer to explicitly start the SRH, breaking the management
isolation provided by the layered architecture.
*/
func (srh ServerRequestHandler) StartListening() {
	lock = make(chan struct{}, 1)

	listener, _ = net.Listen("tcp", srh.GetAddr())
}

func (srh ServerRequestHandler) GetAddr() string {
	return srh.ServerHost + ":" + strconv.Itoa(srh.ServerPort)
}

func (srh ServerRequestHandler) AcceptNewConnection() {
	conn, err := listener.Accept()
	if (err != nil) {
		log.Fatal("Error while accepting connection ", err)
	}

	lock <- struct{}{}
	clientConnection = conn
}

func (srh ServerRequestHandler) Receive() ([]byte, error) {
	clientMsg := make([]byte, 512)
	_, err := clientConnection.Read(clientMsg)

	return clientMsg, err
}

func (srh ServerRequestHandler) Send(msg []byte) {
	_, err := clientConnection.Write(msg)
	if (err != nil) {
		log.Fatal("Error writing response to client. ", err)
	}

	clientConnection.Close()
	<- lock
}

func (srh ServerRequestHandler) StopListening() {
	err := listener.Close()

	if (err != nil) {
		log.Fatal("Error closing listener. ", err)
	}
}