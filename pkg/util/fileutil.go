package util

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadWithIOUtil(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}

func WriteResult(filename string, results []string) {
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		fmt.Printf("创建文件失败: %v", err)
		return
	}
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666)
	writer := bufio.NewWriter(file)
	_, _ = writer.WriteString(strings.Join(results, "\n"))
	_ = writer.Flush()
	fmt.Printf("保存位置: %v\n", filename)
}
