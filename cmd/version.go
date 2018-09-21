package cmd

import (
	"fmt"
	"runtime"

	"bitbucket.org/zkrhm-fdn/fdnservice/buildvars"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long:  `Display version and build information about hellogopher.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s %s\n", buildvars.AppName, buildvars.Version)
		fmt.Printf("  Build date: %s\n", buildvars.BuildDate)
		fmt.Printf("  Built with: %s\n", runtime.Version())
		fmt.Printf("  Author: %s\n", buildvars.Author)
	},
}
