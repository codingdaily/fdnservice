// Package main is just an entry point.
package main

import (
	"fmt"
	"os"

	//watch this one. do not import wrong path, because it's the root command.
	"bitbucket.org/zkrhm-fdn/microsvc-starter/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v\n", err)
		os.Exit(1)
	}
}
