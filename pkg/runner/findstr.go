package runner

import (
	"fmt"
	"go-find/pkg/util"
	"io/fs"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

var (
	matchResults []string
	ch           = make(chan int, 512)
	mutex        sync.Mutex
)

func (e *Engine) Findstr(dir string) []string {
	var pathArray [100]string
	var count = 0
	filepath.Walk(dir, func(curPath string, info fs.FileInfo, err error) error {
		if err != nil {
		}
		if !info.IsDir() {
			// 后缀白名单判断
			ext := path.Ext(curPath)
			if !util.CheckExt(e.FileWhiteExts, ext) {
				return nil
			}
			// 路径黑名单判断
			for _, black := range e.PathBlackWords {
				if strings.Contains(curPath, black) {
					return nil
				}
			}
			// 文件大小判断
			if info.Size() < 1024*1024 {
				pathArray[count] = curPath
				count++
				if count >= 100 {
					count = 0
					go e.findText(pathArray[0:100])
					<-ch
				}
			}
		}
		return nil
	})
	go e.findText(pathArray[0:count])
	<-ch
	return matchResults
}

func (e *Engine) findText(paths []string) {
	for _, curPath := range paths {
		content := util.ReadWithIOUtil(curPath)
		for _, keyword := range e.Keywords {
			index := strings.Index(content, keyword.Keyword)
			if index > -1 {
				var res string
				if keyword.Category != "" {
					res = fmt.Sprintf("%v => [%v] (%v)", curPath, strings.Replace(content[util.Max(0, index-50):util.Min(len(content), index+50)], "\n", "", -1), keyword.Category)
				} else {
					res = fmt.Sprintf("%v => [%v]", curPath, strings.Replace(content[util.Max(0, index-50):util.Min(len(content), index+50)], "\n", "", -1))
				}
				fmt.Printf("[+] %v\n", res)
				mutex.Lock()
				matchResults = append(matchResults, res)
				mutex.Unlock()
				break
			}
		}
	}
	ch <- 1
}
