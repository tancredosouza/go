package common

import (
	"encoding/json"
	"log"
	"../packetdef"
)

type Marshaller struct{}

func (t Marshaller) Marshall(data interface{}) []byte {
	serializedData, err := json.Marshal(data)
	if (err != nil) {
		log.Fatal("Error serializing data. ", err)
	}

	return serializedData
}

func (t Marshaller) Unmarshall(bytes []byte) packetdef.Packet {
	var data packetdef.Packet
	err := json.Unmarshal(bytes, &data)

	if (err != nil) {
		log.Fatal("Error deserializing data. ", err)
	}

	return data
}