package main

import (
	"github.com/spf13/cobra"
	"github.com/vdamery/jdria/cmd/api"
)

var rootCmd *cobra.Command

func init() {
	rootCmd = &cobra.Command{Use: "hr-bootapi-template"}
	rootCmd.AddCommand(api.Start())
}

func main() {
	cobra.CheckErr(rootCmd.Execute())
}
