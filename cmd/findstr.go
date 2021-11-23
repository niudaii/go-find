package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-find/config"
	"go-find/pkg/findstr"
	"go-find/pkg/util"
	"os"
	"strings"
)

func init() {
	findstrCmd.PersistentFlags().StringVarP(&dir, "dir", "d", "", "dir to search")
	findstrCmd.PersistentFlags().StringVarP(&keyword, "keywords", "k", "", "keywords, split by comma")
	findstrCmd.PersistentFlags().StringVarP(&filenameWhiteExt, "fwe", "", "", "filenameWhiteExt, split by comma")
	findstrCmd.PersistentFlags().StringVarP(&pathBlackWord, "pbw", "", "", "pathBlackWord, split by comma")
	findstrCmd.PersistentFlags().StringVarP(&output, "output", "o", "matchResults.txt", "output name")
	rootCmd.AddCommand(findstrCmd)
}

var findstrCmd = &cobra.Command{
	Use:   "findstr",
	Short: fmt.Sprintf("文件名搜索关键词: [%v]\n", strings.Join(config.Conf.ContentKeywords, ",")),
	Run: func(cmd *cobra.Command, args []string) {
		if dir == "" {
			fmt.Println("请输入要搜索的路径")
			os.Exit(1)
		}
		_, err := os.Stat(dir)
		if err != nil {
			fmt.Println("输入的路径不存在")
			os.Exit(1)
		}
		engine := findstr.NewEngine(keyword, filenameWhiteExt, pathBlackWord)
		results := engine.SearchContent(dir)
		fmt.Printf("match count: %v\n", len(results))
		if len(results) > 0 {
			util.WriteResult(output, results)
		}
	},
}
