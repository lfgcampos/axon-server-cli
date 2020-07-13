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
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func setAxonIQAccessTokenHeader(req *http.Request) {
	if viper.IsSet("token") {
		req.Header.Set("AxonIQ-Access-Token", viper.GetString("token"))
	}
}

func setContentTypeApplicationJsonHeader(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
}

func is2xxSuccessful(statusCode int) bool {
	return statusCode >= 200 && statusCode <= 299
}

func printAction(method string, url string) {
	if viper.IsSet("verbose") {
		_, _ = fmt.Fprintf(os.Stderr, "HTTP request: method = '%s', url = '%s'\n", method, url)
	}
}

// GET - Executes a GET on the given URL.
func GET(url string) ([]byte, error) {
	printAction("GET", url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	setAxonIQAccessTokenHeader(req)

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if !is2xxSuccessful(resp.StatusCode) {
		fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

// POST - Executes POST on the given URL, with the given body
func POST(url string, requestBody []byte) ([]byte, error) {
	printAction("POST", url)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	setAxonIQAccessTokenHeader(req)
	setContentTypeApplicationJsonHeader(req)

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if !is2xxSuccessful(resp.StatusCode) {
		fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

// DELETE - Executes DELETE on the given URL
func DELETE(url string) ([]byte, error) {
	printAction("DELETE", url)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	setAxonIQAccessTokenHeader(req)

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if !is2xxSuccessful(resp.StatusCode) {
		fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
