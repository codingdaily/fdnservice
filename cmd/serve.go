package cmd

import (
	"bitbucket.org/zkrhm-fdn/fire-starter/app"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the content",
	Run: func(cmd *cobra.Command, args []string) {
		theApp := app.NewApp()
		theApp.Initialize()

		theApp.Run(":8000")
	},
}
