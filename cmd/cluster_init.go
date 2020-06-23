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
	initContext string
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
	clusterInitCmd.Flags().StringVarP(&initContext, "context", "c", "", "[Optional - Enterprise Edition only] context to register node in")
}

func initCluster(cmd *cobra.Command, args []string) {
	url := buildURL()
	log.Println("calling: " + url)
	responseBody := httpwrapper.POST(url, nil)
	fmt.Printf("%s\n", responseBody)
}

func buildURL() string {
	url := fmt.Sprintf("%s/v1/public", viper.GetString("server"))
	if len(initContext) > 0 {
		url += "?context=" + initContext
	}
	return url
}
