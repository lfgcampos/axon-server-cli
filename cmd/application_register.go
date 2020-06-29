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
	"axon-server-cli/httpwrapper"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type application struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Token       string `json:"token"`
	Roles       []role `json:"roles"`
}

type role struct {
	Context string   `json:"context"`
	Roles   []string `json:"roles"`
}

var applicationRegisterCmd = &cobra.Command{
	Use:     "register",
	Aliases: []string{"r"},
	Short:   "Register an application",
	Long: `Registers an application with specified name. Roles is a comma seperated list of roles per context, where a role per context is the combination of @, e.g. READ@context1,WRITE@context2. If you do not specify the context for the role it will be for context default.
	If you omit the -T option, Axon Server will generate a unique token for you. Applications must use this token to access Axon Server. Note that this token is only returned once, you will not be able to retrieve this token later.`,
	Run: registerApplication,
}

func init() {
	applicationCmd.AddCommand(applicationRegisterCmd)

	applicationRegisterCmd.Flags().StringP("name", "a", "", "*Name of the application")
	applicationRegisterCmd.Flags().StringP("description", "d", "", "[Optional] Description of the application")
	applicationRegisterCmd.Flags().StringSliceP("roles", "r", []string{}, "Roles for the application, use role@context")
	applicationRegisterCmd.Flags().StringP("token", "T", "", "[Optional] Use this token for the app")
	// required flags
	applicationRegisterCmd.MarkFlagRequired("name")
}

func registerApplication(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString("name")
	description, _ := cmd.Flags().GetString("description")
	token, _ := cmd.Flags().GetString("token")
	roles, _ := cmd.Flags().GetStringArray("roles")

	// TODO: where is the best place for validations?
	if len(token) > 0 && len(token) < 16 {
		log.Fatal("Token must be at least 16 characters")
	}

	url := fmt.Sprintf("%s/v1/applications", viper.GetString("server"))
	postBody := buildApplicationJSON(name, description, token, roles)
	log.Printf("calling: %s\n", url)

	responseBody := httpwrapper.POST(url, postBody)
	fmt.Printf("%s\n", responseBody)
}

func buildApplicationJSON(name string, description string, token string, roles []string) []byte {
	application := &application{
		Name:        name,
		Description: description,
		Token:       token,
		Roles:       buildRoles(roles),
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

func buildRoles(roles []string) []role {
	var rolesPerContext = make(map[string][]string)
	for _, roleAndContext := range roles {
		role, context := splitRoleAndContext(roleAndContext)
		rolesPerContext[context] = append(rolesPerContext[context], role)
	}

	var returnValue []role
	for context, roles := range rolesPerContext {
		newRole := role{
			Context: context,
			Roles:   roles,
		}
		returnValue = append(returnValue, newRole)
	}
	return returnValue
}

func splitRoleAndContext(roleAndContext string) (string, string) {
	x := strings.Split(roleAndContext, "@")
	return x[0], x[1]
}
