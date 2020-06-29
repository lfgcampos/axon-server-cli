package utils

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func Print(any interface{}) {
	if viper.GetBool("json") {
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
