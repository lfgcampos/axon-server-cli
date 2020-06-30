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

// constants
const (
	rolePrimary       = "PRIMARY"
	roleActiveBackup  = "ACTIVE_BACKUP"
	rolePassiveBackup = "PASSIVE_BACKUP"
	roleMessagingOnly = "MESSAGING_ONLY"
)

var (
	contextRegister                                   string
	nodes, activeBackup, passiveBackup, messagingOnly []string
)

type nodeAndRole struct {
	Node string `json:"node"`
	Role string `json:"role"`
}

type contextNode struct {
	Context string        `json:"context"`
	Nodes   []string      `json:"nodes"`
	Roles   []nodeAndRole `json:"roles"`
}

var contextRegisterCmd = &cobra.Command{
	Use:     "register",
	Aliases: []string{"r"},
	Short:   "Creates a new context",
	Long:    `Creates a new context, with the specified nodes assigned to it. If you do not specify nodes, all nodes will be assigned to the context.`,
	Run:     createContext,
}

func init() {
	contextCmd.AddCommand(contextRegisterCmd)

	contextRegisterCmd.Flags().StringVarP(&contextRegister, "context", "c", "", "*Name of the context")
	contextRegisterCmd.Flags().StringSliceVarP(&nodes, "nodes", "n", []string{}, "[Enterprise Edition only] primary member nodes for context")
	contextRegisterCmd.Flags().StringSliceVarP(&activeBackup, "activeBackup", "a", []string{}, "[Optional - Enterprise Edition only] active backup member nodes for context")
	contextRegisterCmd.Flags().StringSliceVarP(&passiveBackup, "passiveBackup", "p", []string{}, "[Optional - Enterprise Edition only] passive backup member nodes for context")
	contextRegisterCmd.Flags().StringSliceVarP(&messagingOnly, "messagingOnly", "m", []string{}, "[Optional - Enterprise Edition only] messaging-only member nodes for context")
	// required flags
	contextRegisterCmd.MarkFlagRequired("context")
}

func createContext(cmd *cobra.Command, args []string) {
	url := fmt.Sprintf("%s/v1/context", viper.GetString("server"))
	utils.Print(url)

	contextJSON := buildContextJSON()
	utils.Print(contextJSON)

	responseBody := httpwrapper.POST(url, contextJSON)
	utils.Print(responseBody)
}

func buildContextJSON() []byte {
	var nodesAndRoles []nodeAndRole
	var definedNodes []string
	// build nodes and nodesAndRoles
	definedNodes, nodesAndRoles = addNodes(definedNodes, nodesAndRoles, nodes, rolePrimary)
	definedNodes, nodesAndRoles = addNodes(definedNodes, nodesAndRoles, activeBackup, roleActiveBackup)
	definedNodes, nodesAndRoles = addNodes(definedNodes, nodesAndRoles, passiveBackup, rolePassiveBackup)
	definedNodes, nodesAndRoles = addNodes(definedNodes, nodesAndRoles, messagingOnly, roleMessagingOnly)

	contextNode := &contextNode{
		Context: contextRegister,
		Nodes:   definedNodes,
		Roles:   nodesAndRoles,
	}
	return utils.ToJSON(contextNode)
}

func addNodes(definedNodes []string, nodesAndRoles []nodeAndRole, nodes []string, role string) ([]string, []nodeAndRole) {
	for _, value := range nodes {
		// check if the node is already in use
		if isNodeAlreadyUsed(definedNodes, value) {
			log.Fatal("Node can only be provided once: ", value)
		}
		currentNodeAndRole := nodeAndRole{
			Node: value,
			Role: role,
		}
		nodesAndRoles = append(nodesAndRoles, currentNodeAndRole)
		definedNodes = append(definedNodes, value)
	}
	return definedNodes, nodesAndRoles
}

func isNodeAlreadyUsed(definedNodes []string, node string) bool {
	for _, a := range definedNodes {
		if a == node {
			return true
		}
	}
	return false
}
