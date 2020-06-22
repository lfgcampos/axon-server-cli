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
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"

	"axon-server-cli/utils"
)

var applicationListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List all applications",
	Long:    `Lists all applications and the roles per application per context.`,
	Run:     listApplications,
}

func init() {
	applicationCmd.AddCommand(applicationListCmd)
}

func listApplications(cmd *cobra.Command, args []string) {
	listApplicationsURL := fmt.Sprintf("%s%s", viper.GetString("server"), applicationListURL)
	log.Printf("calling: %s\n", listApplicationsURL)
	body := utils.GET(listApplicationsURL)
	fmt.Printf("%s\n", body)
}
