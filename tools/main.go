package main

import (
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

func rootCmd ()(*cobra.Command){
	cmd := cobra.Command{
		Use: "licotool",
		Short: "開発用ツール",
		Long: "licoの開発用ツールです",
		SilenceUsage: true,
	}

	return &cmd
}


type cmdVars struct{
	Name string
}

func newcmdCmd()(*cobra.Command){
	cmd := cobra.Command{
		Use: "newcmd BaseFile OutFile Name",
		Short: "新しいコマンドを追加",
		Long: "カスタマイズされたCobraのテンプレートを生成してcmd/に追加します",
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			basefile := args[0]
			vars := cmdVars{
				Name: args[2],
			}
			
			t, err := template.ParseFiles(basefile)
			if err!=nil{
				return err
			}

			writeTo, err := os.Create(args[1])
			if err !=nil{
				return err
			}
			defer writeTo.Close()


			if err = t.Execute(writeTo, vars); err!=nil{
				return err
			}

			return nil
		},
	}

	return &cmd
}

func main(){
	root := rootCmd()
	root.AddCommand(newcmdCmd())

	if err:= root.Execute(); err !=nil{
		os.Exit(1)
	}
}
