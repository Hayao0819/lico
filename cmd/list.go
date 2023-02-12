package cmd

import (
	"fmt"
	//"os"

	"github.com/Hayao0819/lico/conf"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "ドットファイルの一覧を表示",
	Long: ``,
	RunE: runList,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func runList(cmd *cobra.Command, args []string)(error){
	list, err := conf.ReadConf(listFile)
	if err !=nil{
		//fmt.Fprintln(os.Stderr, err)
		return err
	}

	for _, entry := range *list{
		parsedRepoPath, _ := conf.Format(string(entry.RepoPath))
		parsedHomePath, _ := conf.Format(string(entry.HomePath))
		fmt.Printf("%v => %v\n" , parsedRepoPath, parsedHomePath)
		
	}
	return nil
}
