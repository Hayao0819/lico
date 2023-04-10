package cmd

import (
	"github.com/Hayao0819/lico/vars"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

func openCmd()(*cobra.Command){
	cmd := cobra.Command{
		Use: "open",
		Short: "設定ディレクトリを開きます",
		Long: "設定ディレクトリをOSのデフォルトのアプリケーションで開きます",
		RunE: func(cmd *cobra.Command, args []string) error {

			open.Run(vars.RepoDir)
			return nil
		},
	}

	return &cmd
}

func init(){
	root.AddCommand(openCmd())
}
