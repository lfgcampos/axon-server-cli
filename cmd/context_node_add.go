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

// contextNodeDeleteCmd represents the DeleteNodeFromContext command
var contextNodeAddCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add a node to a context",
	Long:    `Adds a node with an optional role to the specified context.`,
	Run:     addNodeToContext,
}

func init() {
	contextNodeCmd.AddCommand(contextNodeAddCmd)
	contextNodeAddCmd.Flags().StringP("context", "c", "", "*Name of the context")
	contextNodeAddCmd.Flags().StringP("node", "n", "", "*Name of the node")
	contextNodeAddCmd.Flags().StringP("role", "r", "", "Role of the node (PRIMARY, MESSAGING_ONLY, ACTIVE_BACKUP, PASSIVE_BACKUP)")
	// required flags
	contextNodeAddCmd.MarkFlagRequired("context")
	contextNodeAddCmd.MarkFlagRequired("node")
}

func addNodeToContext(cmd *cobra.Command, args []string) {
	contextName, _ := cmd.Flags().GetString("context")
	nodeName, _ := cmd.Flags().GetString("node")
	nodeRole, _ := cmd.Flags().GetString("role")

	url := buildContextNodeAddURL(contextName, nodeName, nodeRole)
	utils.Print(url)

	responseBody := httpwrapper.POST(url, nil)
	utils.Print(responseBody)
}

func buildContextNodeAddURL(contextName string, nodeName string, nodeRole string) string {
	url := fmt.Sprintf("%s/v1/context/%s/%s", viper.GetString("server"), contextName, nodeName)
	if len(nodeRole) > 0 {
		url += "?role=" + nodeRole
	}
	return url
}
