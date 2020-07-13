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

var clusterInitCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"l"},
	Short:   "initializes the cluster",
	Long:    `Initializes the cluster, creates the _admin context and the specified context. If no context specified it creates default context.`,
	Run:     initCluster,
}

func init() {
	clusterCmd.AddCommand(clusterInitCmd)
	clusterInitCmd.Flags().StringP("context", "c", "", "[Optional - Enterprise Edition only] context to register node in")
}

func initCluster(cmd *cobra.Command, args []string) {
	context, _ := cmd.Flags().GetString("context")
	url := buildContextURL(context)
	utils.Print(url)

	responseBody, err := httpwrapper.GET(url)
	if err != nil {
		log.Fatal(err)
	}
	utils.Print(responseBody)
}

func buildContextURL(context string) string {
	url := fmt.Sprintf("%s/v1/public", viper.GetString("server"))
	if len(context) > 0 {
		url += "?context=" + context
	}
	return url
}
