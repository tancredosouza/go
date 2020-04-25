package common

import (
	"encoding/json"
	"log"
)

type Marshaller struct{}

func (t Marshaller) Marshall(data string) []byte {
	serializedData, err := json.Marshal(data)
	if (err != nil) {
		log.Fatal("Error serializing data. ", err)
	}

	return serializedData
}

func (t Marshaller) Unmarshall(bytes []byte) string {
	var data string
	err := json.Unmarshal(bytes, &data)

	if (err != nil) {
		log.Fatal("Error deserializing data. ", err)
	}

	return data
}