package cmd

import (
	"fmt"
	//"runtime/debug"

	"github.com/Hayao0819/lico/vars"
	"github.com/spf13/cobra"
)

func versionCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:     "version",
		Short:   "バージョン情報",
		Long:    `バージョン情報を表示します`,
		Args:    cobra.NoArgs,
		Aliases: []string{},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Version: %v\n Commit: %v\n   Date: %v\n", vars.Version.Name, vars.Version.Commit, vars.Version.Date)
			return nil
		},
	}

	return &cmd
}

func init() {
	cmd := CmdFunc(versionCmd)
	AddCommand(&cmd)
}
