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
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type user struct {
	Username string   `json:"userName"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
}

var userRegisterCmd = &cobra.Command{
	Use:     "register",
	Aliases: []string{"r"},
	Short:   "Register a user",
	Long: `Registers a user with specified roles. If no roles are specified, Axon Server registers the user with READ role. Specify multiple roles by giving a comma separated list (without spaces), e.g. READ,ADMIN.
If you do not specify a password with the -p option, the command line interface will prompt you for one.`,
	Run: registerUser,
}

func init() {
	userCmd.AddCommand(userRegisterCmd)

	userRegisterCmd.Flags().StringP("username", "u", "", "*Username")
	userRegisterCmd.Flags().StringP("password", "p", "", "[Optional] Password for the user")
	userRegisterCmd.Flags().StringSliceP("roles", "r", []string{}, "[Optional] Roles for the user")
	// required flags
	userRegisterCmd.MarkFlagRequired("username")
}

func registerUser(cmd *cobra.Command, args []string) {
	url := fmt.Sprintf("%s/v1/users", viper.GetString("server"))

	username, _ := cmd.Flags().GetString("username")
	password, _ := cmd.Flags().GetString("password")
	roles, _ := cmd.Flags().GetStringSlice("roles")

	user := &user{
		Username: username,
		Password: password,
		Roles:    roles,
	}
	userJSON := utils.ToJSON(user)
	utils.Print(userJSON)

	responseBody, err := httpwrapper.POST(url, userJSON)
	if err != nil {
		log.Fatal(err)
	}
	utils.Print(responseBody)
}
