package main

import (
	"encoding/json"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

func rootCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:          "licotool",
		Short:        "開発用ツール",
		Long:         "licoの開発用ツールです",
		SilenceUsage: true,
	}

	return &cmd
}

func newcmdCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "newcmd BaseFile OutFile Name",
		Short: "新しいコマンドを追加",
		Long:  "カスタマイズされたCobraのテンプレートを生成してcmd/に追加します",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			basefile := args[0]

			vars := map[string]string{
				"Name": args[2],
			}

			t, err := template.ParseFiles(basefile)
			if err != nil {
				return err
			}

			writeTo, err := os.Create(args[1])
			if err != nil {
				return err
			}
			defer writeTo.Close()

			if err = t.Execute(writeTo, vars); err != nil {
				return err
			}

			return nil
		},
	}

	return &cmd
}

func artifactCmd()*cobra.Command{
	cmd := cobra.Command{
		Use:   "artifact PATH",
		Short: "goreleaserのjsonを解析してパスを返します",
		Long:  "goreleaserのjsonを解析してパスを返します",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			jsonpath := args[0]

			// read json
			jsondata, err := os.ReadFile(jsonpath)
			if err != nil {
				return err
			}

			// parse json
			var artifacts []map[string]interface{}
			json.Unmarshal(jsondata, &artifacts)

			// get path
			cmd.Println(artifacts[0]["path"])

			return nil
		},
	}
	return &cmd
}

func main() {
	root := rootCmd()
	root.AddCommand(newcmdCmd())
	root.AddCommand(artifactCmd())

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
