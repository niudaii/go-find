package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var (
	dir              string
	keyword          string
	pathBlackWord    string
	filenameWhiteExt string
	output           string
)

var rootCmd = &cobra.Command{
	Use:               "go-find",
	Short:             "一款有点好用的文件搜索工具 by zp857",
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
}

func Execute() {
	start := time.Now()
	cobra.CheckErr(rootCmd.Execute())
	fmt.Printf("spent time: %v\n", time.Since(start))
}
