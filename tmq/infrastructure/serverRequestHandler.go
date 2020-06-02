package infrastructure

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

type ServerRequestHandler struct {
	ServerHost string
	ServerPort int
}

var listener net.Listener

/*
As-is, for each connection the code creates a listener and
then closes it when the data is sent. This could harm execution
time. To solve this problem, the function below could be
used instead. However, this would require the Distribution
Layer to explicitly start the SRH, breaking the management
isolation provided by the layered architecture.
*/
func (srh ServerRequestHandler) StartListening() {
	var err error
	listener, err = net.Listen("tcp", srh.GetAddr())
	if (err != nil) {
		log.Fatal("Error while creating listener. ", err)
	}

	log.Println("Listening to new connections!")
}

func (srh ServerRequestHandler) GetAddr() string {
	return srh.ServerHost + ":" + strconv.Itoa(srh.ServerPort)
}

func (srh ServerRequestHandler) KeepAcceptingNewConnections() {
	for {
		srh.AcceptNewConnection()
	}
}

func (srh ServerRequestHandler) AcceptNewConnection() {
		conn, err := listener.Accept()
		if (err != nil) {
			log.Fatal("Error while accepting connection ", err)
		}

		log.Println("Accepted new connection!")
		go handleNewConnection(conn)
}

func handleNewConnection(conn net.Conn) {
	for {
		msgBytes, err := Receive(conn)
		if (err != nil) {
			log.Println("Error receiving message from connection! ", err)
			break;
		}

		text := msgBytes
		log.Println(fmt.Sprintf("Received %s!", text))

		err = Send([]byte(fmt.Sprintf("You've sent %s!", text)), conn)
		if (err != nil) {
			log.Println("Error writing to connection", err)
			break;
		}
	}
}

func CloseConnection(conn net.Conn) {
	err := conn.Close()
	if (err != nil) {
		log.Fatal("Error closing connection! ", err)
	}
}

func (srh ServerRequestHandler) StopListening() {
	err := listener.Close()
	if (err != nil) {
		log.Fatal("Error closing listener. ", err)
	}
}