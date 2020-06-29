package cmd

import (
	"axon-server-cli/utils"
	"github.com/spf13/cobra"
)

type Info struct {
	Version string `json:"Version,omitempty"`
	Commit  string `json:"Commit,omitempty"`
	Date    string `json:"Date,omitempty"`
}

var (
	version    = "dev"
	commit     = "none"
	date       = "unknown"
	versionCmd = &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Version will output the current build information",
		Run:     printVersion,
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

func printVersion(cmd *cobra.Command, args []string) {
	info := Info{
		Version: version,
		Commit:  commit,
		Date:    date,
	}
	utils.Print(info)
}
