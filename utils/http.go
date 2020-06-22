/*
Copyright © 2020 Dusan Perkovic <dusan.perkovic@axoniq.io>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package utils

import (
	"log"
	"time"
	"net/http"
	"io/ioutil"

	"github.com/spf13/viper"
)

const tokenKey = "AxonIQ-Access-Token"

var token string

func init() {
	token = viper.GetString("token")
}

// GET - Executes a GET on the given URL.
func GET(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error reading request ", err)
	}
	req.Header.Set(tokenKey, token)
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}
	return body;
}