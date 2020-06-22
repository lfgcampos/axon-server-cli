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

import "github.com/spf13/cobra"

var (
	clusterListURL           = "/v1/public"
	clusterRegisterNodeURL   = "/v1/cluster"
	clusterUnregisterNodeURL = "/v1/cluster/%s"
	clusterInitURL           = "/v1/context/init"
)

var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Commands related to the cluster",
	Long:  `This is the command related to the cluster`,
}

func init() {
	rootCmd.AddCommand(clusterCmd)
}
