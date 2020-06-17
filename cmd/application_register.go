/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	// "net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type application struct {
	Name string			`json:"name"`
	Description string	`json:"description"`
	Roles []role		`json:"roles"`
}

type role struct {
	Context string	`json:"context"`
	Roles []string	`json:"roles"`
}

var applicationName, applicationDescription string
var applicationRoles []string

var applicationRegisterCmd = &cobra.Command{
	Use:   "register",
	Aliases: []string{"r"},
	Short: "Register an application",
	Long: `use to register an application on the server`,
	Run: registerApplication,
}

func init() {
	applicationCmd.AddCommand(applicationRegisterCmd)

	applicationRegisterCmd.Flags().StringVarP(&applicationName, "name", "a", "", "application name")
	applicationRegisterCmd.Flags().StringVarP(&applicationDescription, "description", "d", "", "application description")
	applicationRegisterCmd.Flags().StringSliceVarP(&applicationRoles, "roles", "r", []string{}, "application roles. Please write them as comma separated values, and separate the context and role with an @ sign. Example: CONTEXT@ROLE1,CONTEXT@ROLE2,...")
}

func registerApplication(cmd *cobra.Command, args []string) {
	applicationURL := viper.GetString("server") + applicationRegisterURL
	log.Printf("calling: %s", applicationURL)
	postBody := buildApplicationJSON()
	req, err := http.NewRequest("POST", applicationURL, bytes.NewBuffer(postBody))
	if err != nil {
		log.Fatal("Error reading request.", err)
	}

	req.Header.Set(axonTokenKey, viper.GetString("token"))
	req.Header.Set(contentType, "application/json")
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

func buildApplicationJSON() []byte {

	application := &application{
		Name: applicationName,
		Description: applicationDescription,
		Roles: 	buildRoles(),
	}
	applicationJSON, err := json.Marshal(&application)
	if err != nil {
		log.Fatal("Error building the application json. ", err)
	}
	fmt.Printf("applicationJson %+v\n", string(applicationJSON))
	return applicationJSON
}

func buildRoles() []role {
	var rolesPerContext = make(map[string][]string)
	for _, roleAndContext := range applicationRoles {
		role, context := splitRoleAndContext(roleAndContext)

		rolesPerContext[context] = append(rolesPerContext[context], role)
	}

	var returnValue []role
	for context, roles := range rolesPerContext {
		newRole := role{
			Context: context,
			Roles: roles,
		}
		returnValue = append(returnValue, newRole)
	}
	return returnValue;
}

func splitRoleAndContext(roleAndContext string) (string, string) {
	x := strings.Split(roleAndContext, "@")
	return x[0], x[1]
}
