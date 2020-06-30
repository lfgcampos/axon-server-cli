package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
)

func Print(any interface{}) {
	if viper.GetBool("pretty-json") {
		asPrettyJson(any)
	} else {
		asJson(any)
	}
}

func asJson(any interface{}) {
	byteArray, isByteArray := any.([]byte)
	if !isByteArray {
		byteArray = ToJSON(any)
	}
	fmt.Println(string(byteArray))
}

func asPrettyJson(any interface{}) {
	byteArray, isByteArray := any.([]byte)
	if !isByteArray {
		byteArray = toPrettyJSON(any)
		fmt.Println(string(byteArray))
	} else {
		// we assume it is already a json, let's just indent it
		pretty := new(bytes.Buffer)
		json.Indent(pretty, byteArray, "", "  ")
		fmt.Println(pretty)
	}
}
