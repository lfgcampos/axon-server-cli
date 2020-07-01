package utils

import (
	"encoding/json"
	"log"
)

func ToJSON(any interface{}) []byte {
	bytes, err := json.Marshal(any)
	if err != nil {
		log.Println("Failed to generate json", err)
	}
	return bytes
}

func toPrettyJSON(any interface{}) []byte {
	bytes, err := json.MarshalIndent(any, "", "  ")
	if err != nil {
		log.Println("Failed to generate json", err)
	}
	return bytes
}
