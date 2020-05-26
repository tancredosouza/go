package infrastructure

import (
	"encoding/binary"
	"log"
	"net"
	"strconv"
)

type ClientRequestHandler struct {
	ServerHost string
	ServerPort int
	receivedBuffer chan []byte
	toSendBuffer chan []byte
	conn net.Conn
}

func (crh ClientRequestHandler) GetAddr() string {
	return crh.ServerHost + ":" + strconv.Itoa(crh.ServerPort)
}

func (crh *ClientRequestHandler) Initialize() {
	crh.receivedBuffer = make(chan []byte, 100)
	crh.toSendBuffer = make(chan []byte, 100)

	var err error
	crh.conn, err = net.Dial("tcp", crh.GetAddr())
	if (err != nil) {
		log.Fatal("Error establishing tcp connection ", err)
	}

	go crh.keepSending()
	go crh.keepReceiving()
}

func (crh *ClientRequestHandler) Send(msgToSend []byte) {
	crh.toSendBuffer <- msgToSend
}

func (crh *ClientRequestHandler) Receive() []byte {
	msgToReceive := <- crh.receivedBuffer

	return msgToReceive
}

func (crh *ClientRequestHandler) keepSending() {
	for {
		msgToSend := <- crh.toSendBuffer
		err := Send(msgToSend, crh.conn)

		if (err != nil) {
			log.Println("error while sending -> ", string(msgToSend), err)
			break
		}
	}
}

func Send(msgToSend []byte, conn net.Conn) error {
	// send message's size
	sizeMsgToServer := make([]byte, 4)
	l := uint32(len(msgToSend))

	binary.LittleEndian.PutUint32(sizeMsgToServer, l)
	conn.Write(sizeMsgToServer)

	// send message
	_, err := conn.Write(msgToSend)

	return err
}

func (crh *ClientRequestHandler) keepReceiving() {
	for {
		receivedData, err := Receive(crh.conn)

		if (err != nil) {
			log.Println("error while receiving -> ", err)
			break
		}

		if (receivedData == nil) {
			continue
		}

		crh.receivedBuffer <- receivedData
	}
}

func Receive(conn net.Conn) ([]byte, error) {
	sizeMsgFromServer := make([]byte, 4)
	n, err := conn.Read(sizeMsgFromServer)

	if (n == 0 || err != nil) {
		return nil, err
	}

	size := binary.LittleEndian.Uint32(sizeMsgFromServer)

	// wait for response and return
	responseMsg := make([]byte, size)
	_, err = conn.Read(responseMsg)

	return responseMsg, err
}

func (crh *ClientRequestHandler) CloseConnection() {
	crh.conn.Close()
}