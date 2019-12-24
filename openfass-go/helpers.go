package openfass_go

import (
	"encoding/json"
	"log"
)

// GetRequestBody of function defincation
func GetRequestBody(data FunctionDefintion) []byte {
	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	return b
}

// Returns the bytes of an interface
func GetByteData(data interface{}) []byte {
	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	return b
}
