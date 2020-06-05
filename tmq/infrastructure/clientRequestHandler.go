package infrastructure

import (
	"errors"
	"log"
	"net"
	"strconv"
	"time"
)

type ClientRequestHandler struct {
	ServerHost string
	ServerPort int
	conn       net.Conn
	ToSendBuffer chan []byte
}

func (crh *ClientRequestHandler) GetAddr() string {
	return crh.ServerHost + ":" + strconv.Itoa(crh.ServerPort)
}

func (crh *ClientRequestHandler) Initialize() {
	crh.StablishConnection()

	crh.ToSendBuffer = make(chan []byte, 100)
	go func() { crh.TryToSend(<- crh.ToSendBuffer) }()

	log.Println("Successfully initialized client connection!")
}

func (crh *ClientRequestHandler) StablishConnection() {
	done := make(chan struct{}, 1)
	go func() {
		for {
			var err error
			crh.conn, err = net.Dial("tcp", crh.GetAddr())
			if err == nil {
				done <- struct{}{}
				break
			}
		}
	}()

	select {
	case <- done:
		log.Println("Connection successfully stablished!")
	case <- time.After(6 * time.Second):
		panic(errors.New("Timeout after 6 seconds waiting for connection to stablish!"))
	}
}

func (crh *ClientRequestHandler) Send(msgToSend []byte) {
	crh.ToSendBuffer <- msgToSend
}

func (crh *ClientRequestHandler) TryToSend(msg []byte) {
	err := Send(msg, crh.conn)
	if (err != nil) {
		log.Println("Tried to send, but couldn't!")
		crh.StablishConnection()
		crh.TryToSend(msg)
	} else {
		log.Println("Sent! ", string(msg))
		crh.TryToSend(<- crh.ToSendBuffer)
	}
}

func (crh *ClientRequestHandler) Receive() []byte {
	for {
		msg, err := Receive(crh.conn)
		if (err == nil) {
			log.Println("Received ", string(msg))
			return msg
		}
	}
}

func (crh *ClientRequestHandler) CloseConnection() {
	crh.conn.Close()
}