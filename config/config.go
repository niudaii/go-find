package config

import (
	"embed"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

//go:embed config.yaml
var f embed.FS

var Conf = new(InitConfig)

type InitConfig struct {
	PathBlackWords    []string `yaml:"path_black_words"`
	FilenameWhiteExts []string `yaml:"filename_white_exts"`
	FilenameKeywords  []string `yaml:"filename_keywords"`
	ContentKeywords   []string `yaml:"content_keywords"`
}

type Keyword struct {
	Keyword  string
	Category string
}

func init() {
	data, _ := f.ReadFile("config.yaml")
	err := yaml.Unmarshal(data, &Conf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
