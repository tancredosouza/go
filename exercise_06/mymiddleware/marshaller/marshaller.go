package marshaller

import (
	"bytes"
	"github.com/my/repo/mymiddleware/protocol"
	"github.com/vmihailenco/msgpack"
	"log"
)

type Marshaller struct{}

func (t Marshaller) Marshall(data interface{}) []byte {
	serializedData, err := msgpack.Marshal(data)
	if (err != nil) {
		log.Fatal("Error serializing data. ", err)
	}

	return serializedData
}

func (t Marshaller) Unmarshall(b []byte) protocol.Packet {
	var data protocol.Packet
	err := msgpack.Unmarshal(bytes.Trim(b, "\x00"), &data)
	if (err != nil) {
		log.Fatal("Error deserializing data. ", err)
	}

	return data
}