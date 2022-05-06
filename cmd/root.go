package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var (
	dir           string
	keyword       string
	pathBlackWord string
	fileWhiteExt  string
	output        string
)

var rootCmd = &cobra.Command{
	Use:               "go-find",
	Short:             "一款有点好用的文件搜索工具 by zp857",
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if dir == "" {
			fmt.Println("请输入要搜索的路径")
			return
		}
		_, err := os.Stat(dir)
		if err != nil {
			fmt.Println("输入的路径不存在")
			return
		}
	},
}

func Execute() {
	rootCmd.PersistentFlags().StringVarP(&dir, "dir", "d", "./", "dir to search")
	rootCmd.PersistentFlags().StringVarP(&keyword, "keywords", "k", "", "keywords, split by comma")
	rootCmd.PersistentFlags().StringVarP(&fileWhiteExt, "file-white-ext", "", "", "file white ext, split by comma")
	rootCmd.PersistentFlags().StringVarP(&pathBlackWord, "path-black-word", "", "", "path black word, split by comma")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "gofind_result.txt", "file to write output[xxx.txt]")
	start := time.Now()
	cobra.CheckErr(rootCmd.Execute())
	fmt.Printf("运行时间: %v\n", time.Since(start))
}
