package runner

import (
	"fmt"
	"go-find/pkg/util"
	"io/fs"
	"path"
	"path/filepath"
	"strings"
)

func (e *Engine) Find(dir string) []string {
	var results []string
	filepath.Walk(dir, func(curPath string, info fs.FileInfo, err error) error {
		if err != nil {
		}
		if !info.IsDir() {
			// 文件名是否包含关键词
			for _, keyword := range e.Keywords {
				if strings.Contains(info.Name(), keyword.Keyword) {
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
					var res string
					if keyword.Category != "" {
						res = fmt.Sprintf("%v (%v)", curPath, keyword.Category)
					} else {
						res = fmt.Sprintf("%v", curPath)
					}
					fmt.Printf("[+] %v\n", res)
					results = append(results, res)
					return nil
				}
			}
		}
		return nil
	})
	return results
}
