package marshaller

import (
	"bytes"
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

func (t Marshaller) Unmarshall(b []byte) string {
	var data string

	err := json.Unmarshal(bytes.Trim(b, "\x00"), &data)

	if (err != nil) {
		log.Println(string(bytes.Trim(b, "\x00")))
		log.Fatal("Error deserializing data. ", err)
	}

	return data
}