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
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"axon-server-cli/utils"
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
	Long: `Registers an application with specified name. Roles is a comma seperated list of roles per context, where a role per context is the combination of @, e.g. READ@context1,WRITE@context2. If you do not specify the context for the role it will be for context default.
	If you omit the -T option, Axon Server will generate a unique token for you. Applications must use this token to access Axon Server. Note that this token is only returned once, you will not be able to retrieve this token later.`,
	Run: registerApplication,
}

func init() {
	applicationCmd.AddCommand(applicationRegisterCmd)

	applicationRegisterCmd.Flags().StringVarP(&applicationName, "name", "a", "", "*Name of the application")
	applicationRegisterCmd.Flags().StringVarP(&applicationDescription, "description", "d", "", "[Optional] Description of the application")
	applicationRegisterCmd.Flags().StringSliceVarP(&applicationRoles, "roles", "r", []string{}, "Roles for the application, use role@context")
	// required flags
	applicationRegisterCmd.MarkFlagRequired("name")
}

func registerApplication(cmd *cobra.Command, args []string) {
	applicationURL := fmt.Sprintf("%s%s", viper.GetString("server"), applicationRegisterURL)
	postBody := buildApplicationJSON()
	log.Printf("calling: %s\n", applicationURL)

	responseBody := utils.POST(applicationURL, postBody)
	fmt.Printf("%s\n", responseBody)
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
	prettyJSON, err := json.MarshalIndent(application, "", "  ")
	if err != nil {
        log.Fatal("Failed to generate json", err)
    }
	fmt.Printf("applicationJson:\n%s\n", string(prettyJSON))
	fmt.Println("applicationJson:")
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
