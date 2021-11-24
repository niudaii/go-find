package find

import (
	"fmt"
	"go-find/config"
	"io/fs"
	"path/filepath"
	"strings"
)

type Engine struct {
	Keywords          []string
	FilenameWhiteExts []string
	PathBlackWords    []string
}

func NewEngine(keyword, filenameWhiteExt, pathBlackWord string) *Engine {
	var keywords []string
	var filenameWhiteExts []string
	var pathBlackWords []string
	if keyword == "" {
		keywords = config.Conf.FilenameKeywords
	} else {
		keywords = strings.Split(keyword, ",")
	}
	if filenameWhiteExt == "" {
		filenameWhiteExts = config.Conf.FilenameWhiteExts
	} else {
		filenameWhiteExts = strings.Split(filenameWhiteExt, ",")
	}
	if pathBlackWord == "" {
		pathBlackWords = config.Conf.PathBlackWords
	} else {
		pathBlackWords = strings.Split(pathBlackWord, ",")
	}
	return &Engine{
		Keywords:          keywords,
		FilenameWhiteExts: filenameWhiteExts,
		PathBlackWords:    pathBlackWords,
	}
}

var (
	matchResults []string
)

func (e *Engine) SearchFilename(dir string) []string {
	fmt.Printf("search filename: [%v]\n", strings.Join(e.Keywords, ","))
	filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
		}
		if !info.IsDir() {
			// 关键词匹配
			for _, keyword := range e.Keywords {
				if strings.Contains(info.Name(), keyword) {
					// 后缀白名单判断
					for _, ext := range e.FilenameWhiteExts {
						if strings.HasSuffix(info.Name(), ext) {
							// 路径黑名单判断
							for _, black := range e.PathBlackWords {
								if strings.Contains(path, black) {
									return nil
								}
							}
							fmt.Println("[+] " + path)
							matchResults = append(matchResults, path)
							break
						}
					}
				}
			}
		}
		return nil
	})
	return matchResults
}
