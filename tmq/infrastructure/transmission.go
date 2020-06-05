package infrastructure

import (
	"encoding/binary"
	"errors"
	"net"
)

func Receive(conn net.Conn) ([]byte, error) {
	msgSize := make([]byte, 4)
	if (conn != nil) {
		n, err := conn.Read(msgSize)

		if (n == 0 || err != nil) {
			return nil, err
		}

		size := binary.LittleEndian.Uint32(msgSize)

		msg := make([]byte, size)
		_, err = conn.Read(msg)
		return msg, err
	} else {
		return nil, errors.New("nil connection")
	}
}

func Send(msg []byte, conn net.Conn) error {
	msgSize := make([]byte, 4)
	l := uint32(len(msg))

	binary.LittleEndian.PutUint32(msgSize, l)

	if (conn != nil) {
		_, err := conn.Write(msgSize)

		if (err != nil) {
			return err
		}
	} else {
		return errors.New("nil connection")
	}

	_, err := conn.Write(msg)
	return err
}