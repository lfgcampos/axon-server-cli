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

var (
	nodename string
)

var clusterUnregisterNodeCmd = &cobra.Command{
	Use:     "unregister",
	Aliases: []string{"u"},
	Short:   "Unregister a node from the cluster",
	Long:    `Removes the node with specified name from the cluster. After this, the node that is deleted will still be running in standalone mode`,
	Run:     unregisterNodeToCluster,
}

func init() {
	clusterCmd.AddCommand(clusterUnregisterNodeCmd)
	// flags
	clusterUnregisterNodeCmd.Flags().StringVarP(&nodename, "node", "n", "", "*Name of the node")
	// required flags
	clusterUnregisterNodeCmd.MarkFlagRequired("node")
}

func unregisterNodeToCluster(cmd *cobra.Command, args []string) {
	url := fmt.Sprintf(clusterUnregisterNodeURL, nodename)
	fullDeleteURL := fmt.Sprintf(viper.GetString("server"), url)
	log.Printf("calling: %s\n", fullDeleteURL)

	responseBody := httpwrapper.DELETE(fullDeleteURL)

	fmt.Printf("%s\n", responseBody)
}
