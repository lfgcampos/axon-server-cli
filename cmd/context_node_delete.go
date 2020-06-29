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
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// contextNodeDeleteCmd represents the DeleteNodeFromContext command
var contextNodeDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short:   "Deletes the node from the context",
	Long:    `Deletes the node with name from the specified context.`,
	Run:     deleteNodeFromContext,
}

func init() {
	contextNodeCmd.AddCommand(contextNodeDeleteCmd)
	contextNodeDeleteCmd.Flags().StringP("context", "c", "", "*Name of the context")
	contextNodeDeleteCmd.Flags().StringP("node", "n", "", "*Name of the node")
	contextNodeDeleteCmd.Flags().BoolP("preserve-event-store", "p", false, "[Optional - Enterprise Edition only] keep event store contents")
	// required flags
	contextNodeDeleteCmd.MarkFlagRequired("context")
	contextNodeDeleteCmd.MarkFlagRequired("node")
}

func deleteNodeFromContext(cmd *cobra.Command, args []string) {
	contextName, _ := cmd.Flags().GetString("context")
	nodeName, _ := cmd.Flags().GetString("node")
	preserveEventStore, _ := cmd.Flags().GetBool("preserve-event-store")

	url := buildContextNodeDeleteURL(contextName, nodeName, preserveEventStore)
	log.Printf("calling: %s\n", url)

	responseBody := httpwrapper.DELETE(url)
	fmt.Printf("%s\n", responseBody)
}

func buildContextNodeDeleteURL(contextName string, nodeName string, preserveEventStore bool) string {
	url := fmt.Sprintf("%s/v1/context/%s/%s", viper.GetString("server"), contextName, nodeName)
	if preserveEventStore {
		url += "?preserveEventStore=true"
	}
	return url
}
