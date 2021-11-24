package findstr

import (
	"fmt"
	"go-find/config"
	"go-find/pkg/util"
	"io/fs"
	"path/filepath"
	"strings"
	"sync"
)

var (
	matchResults []string
	ch           = make(chan int, 512)
	mutex        sync.Mutex
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
		keywordArr = config.Conf.ContentKeywords
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

func (e *Engine) SearchContent(dir string) []string {
	fmt.Printf("search content: [%v]\n", strings.Join(e.KeywordsArr, ","))
	var pathArray [100]string
	var count = 0
	filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
		}
		if !info.IsDir() {
			// 后缀白名单判断
			for _, ext := range e.FilenameWhiteExts {
				if strings.HasSuffix(info.Name(), ext) {
					// 路径黑名单判断
					for _, black := range e.PathBlackWords {
						if strings.Contains(path, black) {
							return nil
						}
					}
					// 文件大小判断
					if info.Size() < 1024*1024 {
						pathArray[count] = path
						count++
						if count >= 100 {
							count = 0
							go e.findText(pathArray[0:100])
							<-ch
						}
					}
					break
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
	for _, path := range paths {
		content := util.ReadWithIOUtil(path)
		for _, keyword := range e.Keywords {
			index := strings.Index(content, keyword.Keyword)
			if index > -1 {
				var res string
				if keyword.Category != "" {
					res = fmt.Sprintf("%v => [%v] (%v)", path, strings.Replace(content[util.Max(0, index-10):util.Min(len(content), index+10)], "\n", "", -1), keyword.Category)
				} else {
					res = fmt.Sprintf("%v => [%v]", path, strings.Replace(content[util.Max(0, index-10):util.Min(len(content), index+10)], "\n", "", -1))
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
