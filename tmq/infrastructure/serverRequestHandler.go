package infrastructure

import (
	"../buffers"
	"log"
	"net"
	"strconv"
)

type ServerRequestHandler struct {
	ServerHost     string
	ServerPort     int
	connectionInfo map[string] net.Conn
}

var listener net.Listener

func (srh *ServerRequestHandler) Initialize() {
	srh.StartListening()
	go srh.KeepAcceptingNewConnections()
	srh.connectionInfo = make(map[string] net.Conn)
}

/*
As-is, for each connection the code creates a listener and
then closes it when the data is sent. This could harm execution
time. To solve this problem, the function below could be
used instead. However, this would require the Distribution
Layer to explicitly start the SRH, breaking the management
isolation provided by the layered architecture.
*/
func (srh *ServerRequestHandler) StartListening() {
	var err error
	listener, err = net.Listen("tcp", srh.GetAddr())
	if (err != nil) {
		log.Fatal("Error while creating listener. ", err)
	}

	log.Println("Listening to new connections!")
}

func (srh *ServerRequestHandler) GetAddr() string {
	return srh.ServerHost + ":" + strconv.Itoa(srh.ServerPort)
}

func (srh *ServerRequestHandler) KeepAcceptingNewConnections() {
	for {
		srh.AcceptNewConnection()
	}
}

func (srh *ServerRequestHandler) AcceptNewConnection() {
		c, err := listener.Accept()
		if (err != nil) {
			log.Fatal("Error while accepting connection ", err)
		}

		log.Println("Accepted new connection!")
		id, err := Receive(c)
		if (err != nil) {
			log.Fatal("error receiving component id ", err)
		}

		srh.connectionInfo[string(id)] = c
		go keepReceivingDataFromConn(c)
}

func keepReceivingDataFromConn(conn net.Conn) {
	for {
		msgBytes, err := Receive(conn)
		if (err != nil) {
			log.Println("Error receiving message from connection! ", err)
			break;
		}

		buffers.IncomingMessages <- msgBytes
	}
}

func (srh *ServerRequestHandler) Send(msg []byte, connId string) {
	err := Send(msg, srh.connectionInfo[connId])
	if (err != nil) {
		log.Println("Error sending message back to connection ", err)
	}
}

func CloseConnection(conn net.Conn) {
	err := conn.Close()
	if (err != nil) {
		log.Fatal("Error closing connection! ", err)
	}
}

func (srh *ServerRequestHandler) StopListening() {
	err := listener.Close()
	if (err != nil) {
		log.Fatal("Error closing listener. ", err)
	}
}