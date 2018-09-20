// Package cmd handles the command-line interface for hellogopher.
package cmd

import (
	"github.com/spf13/cobra"
)

// AppName: blablabla
var (
	AppName   = "YourDamnApp"
	ShortDesc = "Application Boiler Plate"
	LongDesc  = `Please Add Description to your app config`
)

// RootCmd is the root for all hello commands.
var RootCmd = &cobra.Command{
	Use:           AppName,
	Short:         ShortDesc,
	Long:          LongDesc,
	SilenceErrors: true,
}
