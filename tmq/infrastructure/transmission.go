package infrastructure

import (
	"encoding/binary"
	"net"
)

func Receive(conn net.Conn) ([]byte, error) {
	msgSize := make([]byte, 4)
	n, err := conn.Read(msgSize)

	if (n == 0 || err != nil) {
		return nil, err
	}

	size := binary.LittleEndian.Uint32(msgSize)

	msg := make([]byte, size)
	_, err = conn.Read(msg)
	return msg, err
}

func Send(msg []byte, conn net.Conn) error {
	msgSize := make([]byte, 4)
	l := uint32(len(msg))

	binary.LittleEndian.PutUint32(msgSize, l)

	_, err := conn.Write(msgSize)

	if (err != nil) {
		return err
	}

	_, err = conn.Write(msg)
	return err
}