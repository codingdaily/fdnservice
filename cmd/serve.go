package cmd

import (
	"fmt"
	"strings"

	"bitbucket.org/zkrhm-fdn/microsvc-starter/logconfig"
	"bitbucket.org/zkrhm-fdn/microsvc-starter/server"
	"github.com/getsentry/raven-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

var (
	configFile string
	err        error
)

func init() {
	viper.SetEnvPrefix("fdnsvc")
	viper.BindEnv("ravendsn")

	if err != nil {
		panic(err)
	}

	port := ":50051"

	cobra.OnInitialize(initConfig)
	serveCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is )")
	serveCmd.PersistentFlags().StringP("port", "p", port, "port of string, set to 50051 when not defined")
	serveCmd.PersistentFlags().StringP("ravendsn", "r", "", "the raven dsn")

	viper.BindPFlag("ravendsn", serveCmd.PersistentFlags().Lookup("ravendsn"))

	RootCmd.AddCommand(serveCmd)

	raven.SetDSN(serveCmd.PersistentFlags().Lookup("ravendsn").Value.String())

}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile("configFile")
	} else {
		//config file search path :
		// - /etc/${appName}
		// - current working dir
		// viper.SetConfigType("json")
		viper.SetConfigName("config")
		viper.AddConfigPath(fmt.Sprint("/etc/", AppName))
		viper.AddConfigPath(".")
	}

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the content run on port 50051 by default",
	Run: func(cmd *cobra.Command, args []string) {

		bs, err := yaml.Marshal(viper.Sub("logging").AllSettings())
		if err != nil {
			panic(err)
		}

		r := strings.NewReplacer(
			"encoderconfig", "encoderConfig",
			"outputpaths", "outputPaths",
			"erroroutputpaths", "errorOutputPaths",
			"initialfields", "initialFields",
		)
		bs = []byte(r.Replace(string(bs)))

		fmt.Println("serveCmd : checking viper configurations *after replacer \n\n", string(bs))

		logger, err := logconfig.NewZapLogger(bs)
		if err != nil {
			panic(err)
		}

		server := server.NewServerWithLogger(logger)
		theport := cmd.PersistentFlags().Lookup("port").Value.String()

		server.Run(theport)
	},
}
