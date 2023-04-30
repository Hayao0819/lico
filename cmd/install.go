package cmd

import (
	//"fmt"

	pkg "github.com/Hayao0819/lico/pkglist"
	"github.com/spf13/cobra"
	//"github.com/Hayao0819/lico/utils"
	//"github.com/Hayao0819/lico/conf"
	//"github.com/Hayao0819/lico/vars"
)

func installCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "install",
		Short: "パッケージをインストール",
		Long:  `lico-pkgs.jsonに定義されているパッケージをインストールします`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := pkg.ReadList()
			if err != nil {
				return err
			}

			return nil
		},
	}

	return &cmd
}

func init() {
	cmd := CmdFunc(installCmd)
	addCommand(&cmd)
}
