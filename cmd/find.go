package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-find/config"
	"go-find/pkg/runner"
	"go-find/pkg/util"
)

func init() {
	rootCmd.AddCommand(findCmd)
}

var findCmd = &cobra.Command{
	Use:   "find",
	Short: fmt.Sprintf("文件名搜索,内置关键词: %v\n", config.Conf.FindKeywords),
	Run: func(cmd *cobra.Command, args []string) {
		engine := runner.NewEngine(keyword, fileWhiteExt, pathBlackWord)
		results := engine.Find(dir)
		fmt.Printf("匹配数量: %v\n", len(results))
		if len(results) > 0 {
			util.WriteResult(output, results)
		}
	},
}
