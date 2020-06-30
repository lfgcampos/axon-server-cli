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

var (
	usernameRegister, password string
	roles                      []string
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

	userRegisterCmd.Flags().StringVarP(&usernameRegister, "username", "u", "", "*Username")
	userRegisterCmd.Flags().StringVarP(&password, "password", "p", "", "[Optional] Password for the user")
	userRegisterCmd.Flags().StringSliceVarP(&roles, "roles", "r", []string{}, "[Optional] Roles for the user")
	// required flags
	userRegisterCmd.MarkFlagRequired("username")
}

func registerUser(cmd *cobra.Command, args []string) {
	url := fmt.Sprintf("%s/v1/users", viper.GetString("server"))
	utils.Print(url)

	userJSON := buildUserJSON()
	utils.Print(userJSON)

	responseBody := httpwrapper.POST(url, userJSON)
	utils.Print(responseBody)
}

func buildUserJSON() []byte {
	user := &user{
		Username: usernameRegister,
		Password: password,
		Roles:    roles,
	}
	return utils.ToJSON(user)
}
