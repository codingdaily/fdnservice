package cmd

import (
	"bitbucket.org/zkrhm-fdn/microsvc-starter/server"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func init() {
	port := ":50051"
	serveCmd.PersistentFlags().StringP("port", "p", port, "port of string, set to 50051 when not defined")
	RootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the content run on port 50051 by default",
	Run: func(cmd *cobra.Command, args []string) {
		server := server.NewServer()
		logger, _ := zap.NewDevelopment()
		theport := cmd.PersistentFlags().Lookup("port").Value.String()

		logger.Info("port: ", zap.String("port", theport))

		server.Run(theport)
	},
}
