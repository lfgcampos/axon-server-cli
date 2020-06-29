package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

func Print(any interface{}, json bool) {
	if json {
		PrintAsPrettyJson(any)
	} else {
		PrintAsJson(any)
	}
}

func PrintAsJson(any interface{}) {
	bytes, err := json.Marshal(any)
	if err != nil {
		log.Println("Failed to generate json", err)
	}
	fmt.Print(string(bytes))
}

func PrintAsPrettyJson(any interface{}) {
	bytes, err := json.MarshalIndent(any, "", "  ")
	if err != nil {
		log.Println("Failed to generate json", err)
	}
	fmt.Print(string(bytes))
}
