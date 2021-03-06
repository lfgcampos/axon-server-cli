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
	"axon-server-cli/utils"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "axon-server-cli",
	Short: "AxonServer-CLI in GO",
	Long:  `This CLI is used to perform actions on AxonServer`,
}

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

	rootCmd.PersistentFlags().String("config", "axonserver-cli", "[Optional] Config file")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	rootCmd.PersistentFlags().StringP("server", "S", "http://localhost:8024", "Server to send command to")
	viper.BindPFlag("server", rootCmd.PersistentFlags().Lookup("server"))
	rootCmd.PersistentFlags().StringP("token", "t", "", "[Optional] Access token to authenticate at server")
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	rootCmd.PersistentFlags().Bool("pretty-json", false, "If enabled, all outputs will be pretty-json formatted")
	viper.BindPFlag("pretty-json", rootCmd.PersistentFlags().Lookup("pretty-json"))
	rootCmd.PersistentFlags().Bool("verbose", false, "If enabled, more output is produced")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if viper.IsSet("config") {
		// Use config file from the flag.
		viper.SetConfigFile(viper.GetString("config"))
	} else {
		// Search config in current directory with name "axonserver-cli" (without extension).
		var defaultDir = "."
		viper.AddConfigPath(defaultDir)
		// check AXONIQ_HOME for the config as well
		if axoniqHome, ok := os.LookupEnv("AXONIQ_HOME"); ok {
			viper.AddConfigPath(axoniqHome)
		}
		// check USERPROFILE or HOME for .axoniq folder with config as well
		if userprofile, ok := os.LookupEnv("USERPROFILE"); ok {
			viper.AddConfigPath(filepath.Join(userprofile, ".axoniq"))
		}
		if home, ok := os.LookupEnv("HOME"); ok {
			viper.AddConfigPath(filepath.Join(home, ".axoniq"))
		}
		// name of the file
		viper.SetConfigName("axonserver-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		utils.Print("Using config file: " + viper.ConfigFileUsed())
	} else if viper.IsSet("verbose") {
		utils.Print(err.Error())
	}
}
