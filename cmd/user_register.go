/*
Copyright © 2020 Lucas Campos <lucas.campos@axoniq.io>

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
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var userRegisterCmd = &cobra.Command{
	Use:     "register",
	Aliases: []string{"r"},
	Short:   "Register a user",
	Long:    `register a user to be used on axonserver`,
	Run:     registerUser,
}

func init() {
	userCmd.AddCommand(userRegisterCmd)

	userRegisterCmd.Flags().StringVarP(&username, "username", "u", "", "user username")
	userRegisterCmd.Flags().StringVarP(&password, "password", "p", "", "user password")
	userRegisterCmd.Flags().StringSliceVarP(&roles, "roles", "r", []string{}, "user roles")
}

func registerUser(cmd *cobra.Command, args []string) {
	log.Println("calling: " + viper.GetString("server") + registerUserUrl)
	userJson := buildUserJson()
	req, err := http.NewRequest("POST", viper.GetString("server")+registerUserUrl, bytes.NewBuffer(userJson))
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	req.Header.Set(axonTokenKey, viper.GetString("token"))
	req.Header.Set(contentType, jsonType)
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
	fmt.Printf("%s\n", body)
}

func buildUserJson() []byte {
	user := &user{
		username: username,
		password: password,
		roles:    roles,
	}
	userJson, err := json.Marshal(&user)
	if err != nil {
		log.Fatal("Error building the user json. ", err)
	}
	fmt.Printf("userJson: %+v\n", string(userJson))
	return userJson
}
