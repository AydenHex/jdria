package api

import (
	"github.com/spf13/cobra"
	"github.com/vdamery/jdria/internal/api"
)

func Start() *cobra.Command {
	cmds := &cobra.Command{
		Use: "api",
		Run: run,
	}

	_ = cmds.PersistentFlags()
	cmds.AddCommand()

	return cmds
}

func run(cmd *cobra.Command, args []string) {
	api.New().Run()
}
