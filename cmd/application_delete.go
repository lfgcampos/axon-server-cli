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
	applicationDelete string
)

// applicationDeleteCmd represents the applicationDelete command
var applicationDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short:   "Deletes the application",
	Long:    `Deletes the application from Axon Server. Applications will no longer be able to connect to Axon Server using this token.`,
	Run:     deleteApplication,
}

func init() {
	applicationCmd.AddCommand(applicationDeleteCmd)
	applicationDeleteCmd.Flags().StringVarP(&applicationDelete, "application", "a", "", "*Name of the application")
	// required flags
	applicationDeleteCmd.MarkFlagRequired("application")
}

func deleteApplication(cmd *cobra.Command, args []string) {
	url := fmt.Sprintf(applicationDeleteURL, applicationDelete)
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
