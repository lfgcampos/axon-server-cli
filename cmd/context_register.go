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
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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
	Short:   "Register a context",
	Long:    `register a context on axonserver`,
	Run:     createContext,
}

func init() {
	contextCmd.AddCommand(contextRegisterCmd)

	contextRegisterCmd.Flags().StringVarP(&contextRegister, "context", "c", "", "Name of the context")
	contextRegisterCmd.Flags().StringSliceVarP(&nodes, "nodes", "n", []string{}, "[Enterprise Edition only] primary member nodes for context")
	contextRegisterCmd.Flags().StringSliceVarP(&activeBackup, "activeBackup", "a", []string{}, "[Optional - Enterprise Edition only] active backup member nodes for context")
	contextRegisterCmd.Flags().StringSliceVarP(&passiveBackup, "passiveBackup", "p", []string{}, "[Optional - Enterprise Edition only] passive backup member nodes for context")
	contextRegisterCmd.Flags().StringSliceVarP(&messagingOnly, "messagingOnly", "m", []string{}, "[Optional - Enterprise Edition only] messaging-only member nodes for context")
	// required flags
	contextRegisterCmd.MarkFlagRequired("context")
}

func createContext(cmd *cobra.Command, args []string) {
	log.Println("calling: " + viper.GetString("server") + contextRegisterURL)
	userJson := buildContextJson()
	req, err := http.NewRequest("POST", viper.GetString("server")+contextRegisterURL, bytes.NewBuffer(userJson))
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	req.Header.Set(axonTokenKey, viper.GetString("token"))
	req.Header.Set(contentType, jsonType)
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}
	fmt.Printf("%s\n", body)
}

func buildContextJson() []byte {
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
	contextNodeJson, err := json.Marshal(&contextNode)
	if err != nil {
		log.Fatal("Error building the contextNode json. ", err)
	}
	fmt.Printf("contextNodeJson: %+v\n", string(contextNodeJson))
	return contextNodeJson
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
