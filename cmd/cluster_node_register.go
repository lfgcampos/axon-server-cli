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

var (
	internalHost      string
	internalPort      string
	contextToRegister string
	noContext         bool
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
	Short:   "register a node to the clusters",
	Long: `Tells this Axon Server node to join a cluster. You need to specify the hostname and optionally the internal port number of a node that is the leader of the _admin context. If you do not specify the internal port number it uses 8224 (default internal port number for Axon Server).
	If you specify a context, the new node will be a member of the specified context. If you haven't specified a context, the new node will become a member of all the contexts which the _admin leader is a part of.`,
	Run: registerNodeToCluster,
}

func init() {
	clusterCmd.AddCommand(clusterRegisterNodeCmd)
	// flags
	clusterRegisterNodeCmd.Flags().StringVarP(&internalHost, "internal-host", "i", "", "Internal hostname of the node")
	clusterRegisterNodeCmd.Flags().StringVarP(&internalPort, "internal-port", "p", "8224", "Internal port of the node (default 8224)")
	clusterRegisterNodeCmd.Flags().StringVarP(&contextToRegister, "context", "c", "", "[Optional - Enterprise Edition only] context to register node in")
	clusterRegisterNodeCmd.Flags().BoolVarP(&noContext, "no-contexts", "", false, "[Optional - Enterprise Edition only] add node to cluster, but don't register it in any contexts")
}

func registerNodeToCluster(cmd *cobra.Command, args []string) {

	if len(contextToRegister) > 0 && noContext {
		log.Fatal("Cannot specify a context when also using 'no-context' option.")
	}

	log.Println("calling: " + viper.GetString("server") + clusterRegisterNodeURL)
	clusterNodeJson := buildClusterNodeJson()
	req, err := http.NewRequest("POST", viper.GetString("server")+clusterRegisterNodeURL, bytes.NewBuffer(clusterNodeJson))
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

func buildClusterNodeJson() []byte {
	clusterNode := &clusterNode{
		InternalHostName: internalHost,
		InternalGrpcPort: internalPort,
		Context:          contextToRegister,
		NoContexts:       noContext,
	}
	clusterNodeJson, err := json.Marshal(&clusterNode)
	if err != nil {
		log.Fatal("Error building the clusterNode json. ", err)
	}
	fmt.Printf("clusterNodeJson: %+v\n", string(clusterNodeJson))
	return clusterNodeJson
}
