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
	"os"

	"github.com/spf13/viper"
)

var (
	// used for flags
	cfgFile  string
	server   string
	token    string
	jsonFlag bool

	rootCmd = &cobra.Command{
		Use:   "axon-server-cli",
		Short: "AxonServer-CLI in GO",
		Long:  `This CLI is used to perform actions on AxonServer`,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is axonserver-cli.yaml)")

	rootCmd.PersistentFlags().BoolVarP(&jsonFlag, "json", "", false, "If enabled, all outputs will be json formatted")
	viper.BindPFlag("json", rootCmd.PersistentFlags().Lookup("json"))
	rootCmd.PersistentFlags().StringVarP(&server, "server", "S", "http://localhost:8024", "Server to send command to")
	viper.BindPFlag("server", rootCmd.PersistentFlags().Lookup("server"))
	rootCmd.PersistentFlags().StringVarP(&token, "access-token", "t", "", "[Optional] Access token to authenticate at server")
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in current directory with name "axonserver-cli" (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigName("axonserver-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("Failed to load config file:", err)
	}
}
