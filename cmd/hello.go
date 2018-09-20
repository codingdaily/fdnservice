package cmd

import (
	"os"

	"bitbucket.org/zkrhm-fdn/fire-starter/hello"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(helloCmd)
}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Say hello",
	Long:  `Print a nice hello message on the standard output.`,
	Run: func(cmd *cobra.Command, args []string) {
		hello.Hello(os.Stdout)
	},
}
