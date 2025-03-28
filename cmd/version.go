package cmd

import (
	"fmt"
	"github.com/manicar2093/gomancer/versioning"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of this tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Your Gomancer version is:", versioning.Version)
	},
}
