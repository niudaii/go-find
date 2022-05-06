package runner

import (
	"fmt"
	"go-find/config"
	"strings"
)

type Engine struct {
	Keywords       []*config.Keyword
	FileWhiteExts  []string
	PathBlackWords []string
}

func NewEngine(keyword, fileWhiteExt, pathBlackWord string) *Engine {
	var keywords []*config.Keyword
	var keywordArr []string
	var fileWhiteExts []string
	var pathBlackWords []string
	if keyword == "" {
		keywordArr = config.Conf.FindKeywords
	} else {
		keywordArr = strings.Split(keyword, ",")
	}
	if fileWhiteExt == "" {
		fileWhiteExts = config.Conf.FileWhiteExts
	} else {
		fileWhiteExts = strings.Split(fileWhiteExt, ",")
	}
	if pathBlackWord == "" {
		pathBlackWords = config.Conf.PathBlackWords
	} else {
		pathBlackWords = strings.Split(pathBlackWord, ",")
	}
	fmt.Printf("path_black_words: %v\n", pathBlackWords)
	fmt.Printf("file_white_exts: %v\n", fileWhiteExts)
	fmt.Printf("keywords: %v\n", keywordArr)
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
		Keywords:       keywords,
		FileWhiteExts:  fileWhiteExts,
		PathBlackWords: pathBlackWords,
	}
}
