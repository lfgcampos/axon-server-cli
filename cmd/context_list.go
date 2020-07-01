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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"axon-server-cli/httpwrapper"
)

var contextListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List all contexts",
	Long:    `Lists all contexts and the nodes assigned to the contexts. Per context it shows the master (responsible for replicating events) and the coordinator (responsible for rebalancing).`,
	Run:     listContexts,
}

func init() {
	contextCmd.AddCommand(contextListCmd)
}

func listContexts(cmd *cobra.Command, args []string) {
	url := fmt.Sprintf("%s/v1/public/context", viper.GetString("server"))
	utils.Print(url)

	responseBody := httpwrapper.GET(url)
	utils.Print(responseBody)
}
