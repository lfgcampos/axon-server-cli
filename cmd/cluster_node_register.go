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

type clusterNode struct {
	InternalHostName string `json:"internalHostName"`
	InternalGrpcPort string `json:"internalGrpcPort"`
	Context          string `json:"context"`
	NoContexts       bool   `json:"noContexts"`
}

var clusterRegisterNodeCmd = &cobra.Command{
	Use:     "register",
	Aliases: []string{"r"},
	Short:   "Register a node to the cluster",
	Long: `Tells this Axon Server node to join a cluster. You need to specify the hostname and optionally the internal port number of a node that is the leader of the _admin context. If you do not specify the internal port number it uses 8224 (default internal port number for Axon Server).
	If you specify a context, the new node will be a member of the specified context. If you haven't specified a context, the new node will become a member of all the contexts which the _admin leader is a part of.`,
	Run: registerNodeToCluster,
}

func init() {
	clusterCmd.AddCommand(clusterRegisterNodeCmd)
	// flags
	clusterRegisterNodeCmd.Flags().StringP("internal-host", "i", "", "*Internal hostname of the node")
	clusterRegisterNodeCmd.Flags().StringP("internal-port", "p", "8224", "Internal port of the node")
	clusterRegisterNodeCmd.Flags().StringP("context", "c", "", "[Optional - Enterprise Edition only] context to register node in")
	clusterRegisterNodeCmd.Flags().BoolP("no-contexts", "", false, "[Optional - Enterprise Edition only] add node to cluster, but don't register it in any contexts")
	// required flags
	clusterRegisterNodeCmd.MarkFlagRequired("internal-host")
}

func registerNodeToCluster(cmd *cobra.Command, args []string) {

	context, _ := cmd.Flags().GetString("context")
	noContexts, _ := cmd.Flags().GetBool("no-contexts")

	if len(context) > 0 && noContexts {
		log.Fatal("Cannot specify a context when also using 'no-context' option.")
	}

	url := fmt.Sprintf("%s/v1/cluster", viper.GetString("server"))

	internalHost, _ := cmd.Flags().GetString("internal-host")
	internalPort, _ := cmd.Flags().GetString("internal-port")

	clusterNode := &clusterNode{
		InternalHostName: internalHost,
		InternalGrpcPort: internalPort,
		Context:          context,
		NoContexts:       noContexts,
	}
	clusterNodeJSON := utils.ToJSON(clusterNode)
	utils.Print(clusterNodeJSON)

	responseBody, err := httpwrapper.POST(url, clusterNodeJSON)
	if err != nil {
		log.Fatal(err)
	}
	utils.Print(responseBody)
}
