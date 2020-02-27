package gofaas

import (
	"encoding/json"
	"log"
)

// GetRequestBody of function definition
func GetRequestBody(data *FunctionDefintion) []byte {
	b, err := json.Marshal(*data)
	if err != nil {
		log.Println(err)
	}
	return b
}

// GetByteData will returns the bytes of an interface
func GetByteData(data interface{}) []byte {
	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	return b
}
