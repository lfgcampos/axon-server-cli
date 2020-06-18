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

import "github.com/spf13/cobra"

var userListURL = "/v1/public/users"
var userRegisterURL = "/v1/users"
var userDeleteURL = "/v1/users/" //{username}

var userCmd = &cobra.Command{
	Use:     "user",
	Aliases: []string{"u"},
	Short:   "Commands related to users",
	Long:    `This is the command related to users`,
}

func init() {
	rootCmd.AddCommand(userCmd)
}
