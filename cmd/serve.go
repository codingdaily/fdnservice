package cmd

import (
	"bitbucket.org/zkrhm-fdn/microsvc-starter/server"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the content",
	Run: func(cmd *cobra.Command, args []string) {
		server := server.NewServer()

		server.Run(":50051")
	},
}
