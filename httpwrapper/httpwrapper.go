/*
Package httpwrapper contains a small wrapper for go's native http methods.

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
package httpwrapper

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"

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

	req.Header.Set("AxonIQ-Access-Token", token)

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body. ", err)
	}

	return responseBody
}

// POST - Executes POST on the given URL, with the given body
func POST(url string, requestBody []byte) []byte {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal("Error reading request.", err)
	}

	req.Header.Set("AxonIQ-Access-Token", token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body. ", err)
	}

	return responseBody
}

// DELETE - Executes DELETE on the given URL, with the given body.
func DELETE(url string) []byte {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatal("Error reading the request. ", err)
	}

	req.Header.Set("AxonIQ-Access-Token", token)
	req.Header.Set("Content-Type", "application/json")
	
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body. ", err)
	}

	return responseBody
}
