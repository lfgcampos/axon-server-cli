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
	"axon-server-cli/utils"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"axon-server-cli/httpwrapper"
)

var userListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List all users",
	Long:    `Returns a list of all registered users and their roles.`,
	Run:     listUsers,
}

func init() {
	userCmd.AddCommand(userListCmd)
}

func listUsers(cmd *cobra.Command, args []string) {
	url := fmt.Sprintf("%s/v1/public/users", viper.GetString("server"))
	utils.Print(url)

	responseBody, err := httpwrapper.GET(url)
	if err != nil {
		log.Fatal(err)
	}
	utils.Print(responseBody)
}
