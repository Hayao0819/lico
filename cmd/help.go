package cmd

import (
	//"fmt"

	"github.com/spf13/cobra"
	//"github.com/Hayao0819/lico/utils"
	//"github.com/Hayao0819/lico/conf"
	//"github.com/Hayao0819/lico/vars"
)

func helpCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "help",
		Short: "ヘルプを表示します",
		//Long: ``,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			root.Help()
			return nil
		},
	}

	return &cmd
}

func init() {
	//root.AddCommand(helpCmd())
	root.SetHelpCommand(helpCmd())
}
