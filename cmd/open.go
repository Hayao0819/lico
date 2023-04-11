package cmd

import (
	"strings"

	"github.com/Hayao0819/lico/utils"
	"github.com/Hayao0819/lico/vars"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

func openCmd()(*cobra.Command){
	cmd := cobra.Command{
		Use: "open COMMAND",
		Short: "設定ディレクトリを開きます",
		Long: "設定ディレクトリをOSのデフォルトのアプリケーションで開きます",
		Args: cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			// コマンド指定なし
			if len(args) ==0{
				return open.Run(vars.RepoDir)
			}

			runcmd := []string{}
			replaced := false
			for _, a := range args{
				if strings.Contains(a, "%s"){
					replaced=true
					runcmd = append(runcmd, strings.ReplaceAll(a, "%s", vars.RepoDir))
				}else{
					runcmd = append(runcmd, a)
				}
			}

			// %sがない場合は末尾に追加する
			if ! replaced{
				runcmd = append(runcmd, vars.RepoDir)
			}

			return utils.RunCmd(runcmd[0], runcmd[1:]...)
		},
	}

	return &cmd
}

func init(){
	root.AddCommand(openCmd())
}
