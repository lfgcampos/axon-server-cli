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
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	nodename string
)

var clusterUnregisterNodeCmd = &cobra.Command{
	Use:     "unregister",
	Aliases: []string{"u"},
	Short:   "unregister a node to the clusters",
	Long:    `Removes the node with specified name from the cluster. After this, the node that is deleted will still be running in standalone mode`,
	Run:     unregisterNodeToCluster,
}

func init() {
	clusterCmd.AddCommand(clusterUnregisterNodeCmd)
	// flags
	clusterUnregisterNodeCmd.Flags().StringVarP(&nodename, "node", "n", "", "Name of the node")
}

func unregisterNodeToCluster(cmd *cobra.Command, args []string) {
	url := fmt.Sprintf(clusterUnregisterNodeURL, nodename)
	log.Println("calling: " + viper.GetString("server") + url)
	req, err := http.NewRequest("DELETE", viper.GetString("server")+url, nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	req.Header.Set(axonTokenKey, viper.GetString("token"))
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
