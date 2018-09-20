// Package main is just an entry point.
package main

import (
	"fmt"
	"os"

	"bitbucket.org/zkrhm-fdn/fire-starter/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v\n", err)
		os.Exit(1)
	}
}
