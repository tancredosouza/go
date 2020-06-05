package infrastructure

import (
	"../buffers"
	"log"
	"../locks"
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
	srh.connectionInfo = make(map[string] net.Conn)
	go srh.KeepAcceptingNewConnections()
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
		b, err := Receive(c)
		if (err != nil) {
			log.Fatal("error receiving component id ", err)
		}

		id := string(b)
		locks.ConnectionInfoLock.Lock()
		srh.connectionInfo[id] = c
		locks.ConnectionInfoLock.Unlock()

		locks.OutgoingLock.Lock()
		if _, ok := buffers.OutgoingMessages[id]; !ok {
			buffers.OutgoingMessages[id] = make([][]byte, 0)
		}
		locks.OutgoingLock.Unlock()

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

func (srh *ServerRequestHandler) Send(connId string) {
	locks.ConnectionInfoLock.Lock()
	if (srh.isValidConnection(connId) && len(buffers.OutgoingMessages[connId]) > 0) {
		msg := buffers.OutgoingMessages[connId][0]
		err := Send(msg, srh.connectionInfo[connId])
		if (err != nil) {
			log.Println("Error sending message back to connection ", err)
			delete(srh.connectionInfo, connId)
		} else {
			buffers.OutgoingMessages[connId] = buffers.OutgoingMessages[connId][1:]
		}
	}
	locks.ConnectionInfoLock.Unlock()
}

func (srh *ServerRequestHandler) isValidConnection(connId string) bool {
	_, ok := srh.connectionInfo[connId];

	return ok
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