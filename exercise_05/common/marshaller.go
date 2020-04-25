package common

import (
	"encoding/json"
	"log"
)

type Marshaller struct{}

func (t Marshaller) Marshall(data interface{}) []byte {
	serializedData, err := json.Marshal(data)
	if (err != nil) {
		log.Fatal("Error serializing data. ", err)
	}

	return serializedData
}

func (t Marshaller) Unmarshall(bytes []byte) interface{} {
	var data interface{}
	err := json.Unmarshal(bytes, &data)

	if (err != nil) {
		log.Fatal("Error deserializing data. ", err)
	}

	return data
}