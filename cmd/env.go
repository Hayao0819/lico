package cmd

import (
	"strings"

	"github.com/Hayao0819/lico/conf"
	"github.com/Hayao0819/lico/osenv"
	"github.com/spf13/cobra"
)

func envCmd() *cobra.Command {
	var showOnlyKey bool

	cmd := cobra.Command{
		Use:   "env [Name]",
		Short: "リストで利用できる変数を表示",
		Long: `リストで利用できる変数を表示します。
リストはGolangのテンプレート構文をサポートしており、環境依存の値を埋め込むことが可能です。
このコマンドではテンプレート構文内で利用可能な変数の一覧を表示します。
`,
		Args: cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// envを取得
			env, err := osenv.Get()
			if err != nil {
				return err
			}

			keys := env.GetSortedKeys()

			if showOnlyKey {
				for _, key := range keys {
					cmd.Println(key)
				}
			} else if len(args) == 0 {
				for _, key := range keys {
					if strings.Contains(env[key], " ") || strings.Contains(env[key], "　") {
						cmd.Printf("%v = \"%v\"\n", key, env[key])
					} else {
						cmd.Printf("%v = %v\n", key, env[key])
					}
				}
			} else if len(args) == 1 {
				for index, value := range env {
					if index == args[0] {
						cmd.Printf("%v\n", value)
					}
				}
			}

			return nil
		},
	}

	cmd.AddCommand(funclistCmd())
	cmd.Flags().BoolVarP(&showOnlyKey, "key", "k", false, "変数名のみを表示")

	return &cmd
}

func funclistCmd ()*cobra.Command{
	cmd := cobra.Command{
		Use: "func",
		Short: "テンプレートで利用可能な関数を表示",
		Long: `テンプレートで利用可能な関数を表示します。`,
		RunE: func(cmd *cobra.Command, args []string) error {
			funcs := conf.GetTemplateFuncMap()
			for key := range *funcs {
				cmd.Println(key)
			}
			return nil
		},
	}

	return &cmd
}

func init() {
	cmd := CmdFunc(envCmd)
	addCommand(&cmd)
}
