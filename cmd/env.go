package cmd

import (
	"fmt"

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
					fmt.Println(key)
				}
			} else if len(args) == 0 {
				for _, key := range keys {
					fmt.Printf("%v = %v\n", key, env[key])
				}
			} else if len(args) == 1 {
				for index, value := range env {
					if index == args[0] {
						fmt.Printf("%v\n", value)
					}
				}
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&showOnlyKey, "key", "k", false, "変数名のみを表示")

	return &cmd
}

func init() {
	root.AddCommand(envCmd())
}
