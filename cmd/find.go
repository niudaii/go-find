package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-find/config"
	"go-find/pkg/find"
	"go-find/pkg/util"
	"os"
	"strings"
)

func init() {
	findCmd.PersistentFlags().StringVarP(&dir, "dir", "d", "", "dir to search")
	findCmd.PersistentFlags().StringVarP(&keyword, "keywords", "k", "", "keywords, split by comma")
	findCmd.PersistentFlags().StringVarP(&filenameWhiteExt, "filename-white-ext", "", "", "filename white ext, split by comma")
	findCmd.PersistentFlags().StringVarP(&pathBlackWord, "path-black-word", "", "", "path black word, split by comma")
	findCmd.PersistentFlags().StringVarP(&output, "output", "o", "find_results.txt", "output name")
	rootCmd.AddCommand(findCmd)
}

var findCmd = &cobra.Command{
	Use:   "find",
	Short: fmt.Sprintf("文件名搜索关键词: [%v]\n", strings.Join(config.Conf.FilenameKeywords, ",")),
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
		engine := find.NewEngine(keyword, filenameWhiteExt, pathBlackWord)
		results := engine.SearchFilename(dir)
		fmt.Printf("match count: %v\n", len(results))
		if len(results) > 0 {
			util.WriteResult(output, results)
		}
	},
}
