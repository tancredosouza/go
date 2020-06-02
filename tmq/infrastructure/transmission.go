package infrastructure

import "net"

func Receive(conn net.Conn) ([]byte, error) {
	msg := make([]byte, 512)
	_, err := conn.Read(msg)
	return msg, err
}

func Send(msg []byte, conn net.Conn) error {
	_, err := conn.Write(msg)
	return err
}