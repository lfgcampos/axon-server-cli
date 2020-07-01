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

// applicationDeleteCmd represents the applicationDelete command
var applicationDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short:   "Deletes the application",
	Long:    `Deletes the application from Axon Server. Applications will no longer be able to connect to Axon Server using this token.`,
	Run:     deleteApplication,
}

func init() {
	applicationCmd.AddCommand(applicationDeleteCmd)
	applicationDeleteCmd.Flags().StringP("application", "a", "", "*Name of the application")
	// required flags
	applicationDeleteCmd.MarkFlagRequired("application")
}

func deleteApplication(cmd *cobra.Command, args []string) {
	applicationName, _ := cmd.Flags().GetString("application")

	url := fmt.Sprintf("%s/v1/applications/%s", viper.GetString("server"), applicationName)
	utils.Print(url)

	responseBody := httpwrapper.DELETE(url)
	utils.Print(responseBody)
}
