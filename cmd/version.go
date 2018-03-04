package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var releaseVersion = "0.1"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gsearch",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gsearch version", releaseVersion, "with", runtime.Version())
	},
}
