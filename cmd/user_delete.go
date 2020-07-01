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
	"axon-server-cli/httpwrapper"
	"axon-server-cli/utils"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// userDeleteCmd represents the deleteUser command
var userDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short:   "Remove the user",
	Long:    `Deletes the specified user from axonserver.`,
	Run:     deleteUser,
}

func init() {
	userCmd.AddCommand(userDeleteCmd)
	userDeleteCmd.Flags().StringP("username", "u", "", "*user username")
	// required flags
	userDeleteCmd.MarkFlagRequired("username")
}

func deleteUser(cmd *cobra.Command, args []string) {
	username, _ := cmd.Flags().GetString("username")
	url := fmt.Sprintf("%s/v1/users/%s", viper.GetString("server"), username)
	utils.Print(url)

	responseBody := httpwrapper.DELETE(url)
	utils.Print(responseBody)
}
