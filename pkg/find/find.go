package find

import (
	"fmt"
	"go-find/config"
	"io/fs"
	"path/filepath"
	"strings"
)

type Engine struct {
	KeywordsArr       []string
	Keywords          []*config.Keyword
	FilenameWhiteExts []string
	PathBlackWords    []string
}

func NewEngine(keyword, filenameWhiteExt, pathBlackWord string) *Engine {
	var keywords []*config.Keyword
	var keywordArr []string
	var filenameWhiteExts []string
	var pathBlackWords []string
	if keyword == "" {
		keywordArr = config.Conf.FilenameKeywords
	} else {
		keywordArr = strings.Split(keyword, ",")
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
	for _, temp := range keywordArr {
		var k string
		var v string
		if strings.Contains(temp, "|") {
			k = strings.Split(temp, "|")[0]
			v = strings.Split(temp, "|")[1]
		} else {
			k = temp
		}
		keywords = append(keywords, &config.Keyword{
			Keyword:  k,
			Category: v,
		})
	}
	return &Engine{
		KeywordsArr:       keywordArr,
		Keywords:          keywords,
		FilenameWhiteExts: filenameWhiteExts,
		PathBlackWords:    pathBlackWords,
	}
}

func (e *Engine) SearchFilename(dir string) []string {
	var matchResults []string
	fmt.Printf("search filename: [%v]\n", strings.Join(e.KeywordsArr, ","))
	filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
		}
		if !info.IsDir() {
			// 关键词匹配
			for _, keyword := range e.Keywords {
				if strings.Contains(info.Name(), keyword.Keyword) {
					// 后缀白名单判断
					for _, ext := range e.FilenameWhiteExts {
						if strings.HasSuffix(info.Name(), ext) {
							// 路径黑名单判断
							for _, black := range e.PathBlackWords {
								if strings.Contains(path, black) {
									return nil
								}
							}
							var res string
							if keyword.Category != "" {
								res = fmt.Sprintf("%v (%v)", path, keyword.Category)
							} else {
								res = fmt.Sprintf("%v", path)
							}
							fmt.Printf("[+] %v\n", res)
							matchResults = append(matchResults, res)
							return nil
						}
					}
				}
			}
		}
		return nil
	})
	return matchResults
}
