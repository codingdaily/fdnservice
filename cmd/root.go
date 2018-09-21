// Package cmd handles the command-line interface for hellogopher.
package cmd

import (
	"bitbucket.org/zkrhm-fdn/fdnservice/buildvars"
	"github.com/spf13/cobra"
)

// AppName: blablabla
var (
	AppName   = "Microservice"
	ShortDesc = "Microservice Boiler Plate"
	LongDesc  = `Please Add Description to your app config`
)

// RootCmd is the root for all hello commands.
var RootCmd = &cobra.Command{
	Use:           buildvars.AppName,
	Short:         buildvars.ShortDesc,
	Long:          buildvars.LongDesc,
	SilenceErrors: true,
}
