package cmd

import (
	"fmt"
	"strings"

	"bitbucket.org/zkrhm-fdn/fdnservice/buildvars"
	"bitbucket.org/zkrhm-fdn/fdnservice/logconfig"
	"bitbucket.org/zkrhm-fdn/fdnservice/server"
	raven "github.com/getsentry/raven-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

var (
	configFile string
	err        error
)

func init() {
	viper.SetEnvPrefix(buildvars.AppName)
	viper.BindEnv("raven_dsn")

	if err != nil {
		panic(err)
	}

	port := ":50051"

	cobra.OnInitialize(initConfig)
	serveCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is )")
	serveCmd.PersistentFlags().StringP("port", "p", port, "port of string, set to 50051 when not defined")
	serveCmd.PersistentFlags().StringP("raven-dsn", "r", "", `the raven dsn, also configurable 
																through environment variable by setting FDNSVC_RAVEN_DSN`)

	viper.BindPFlag("raven_dsn", serveCmd.PersistentFlags().Lookup("raven-dsn"))
	viper.BindPFlag("port", serveCmd.PersistentFlags().Lookup("port"))

	RootCmd.AddCommand(serveCmd)
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
		viper.AddConfigPath(fmt.Sprint("/etc/", buildvars.AppName))
		viper.AddConfigPath(".")
	}

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if viper.GetString("raven_dsn") == "" {
		viper.Set("raven_dsn", viper.GetString("raven.dsn"))
	}

	if viper.GetString("port") == "" {
		viper.Set("port", viper.GetString("app.port"))
	}

	ravenDSN := viper.GetString("raven_dsn")
	raven.SetDSN(ravenDSN)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the content run on port 50051 by default (if not defined on flags, )",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("sub logging : ", viper.Sub("logging").AllSettings())

		bs, err := yaml.Marshal(viper.Sub("logging").AllSettings())
		if err != nil {
			panic(err)
		}
		/* handling viper's "feature"
		viper lowercase all keys, while zap configuration is case sensitive
		so this part of code is written
		*/
		r := strings.NewReplacer(
			"encoderconfig", "encoderConfig",
			"outputpaths", "outputPaths",
			"erroroutputpaths", "errorOutputPaths",
			"initialfields", "initialFields",
			"levelencoder", "levelEncoder",
			"levelkey", "levelKey",
			"messagekey", "messageKey",
		)
		bs = []byte(r.Replace(string(bs)))

		// fmt.Println("serveCmd : checking viper configurations *after replacer \n\n", string(bs))
		raven.CapturePanicAndWait(func() {
			logger, err := logconfig.NewZapLogger(bs)
			defer logger.Sync()
			if err != nil {
				panic(err)
			}
			// panic(errors.New("fuck you biatch! "))

			server := server.NewServerWithLogger(logger)
			theport := cmd.PersistentFlags().Lookup("port").Value.String()

			server.Run(theport)
		}, map[string]string{
			"stage": "creating logger, server and running the server",
		})
	},
}
