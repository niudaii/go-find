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
	PathBlackWords  []string `yaml:"path_black_words"`
	FileWhiteExts   []string `yaml:"file_white_exts"`
	FindKeywords    []string `yaml:"find_keywords"`
	FindstrKeywords []string `yaml:"findstr_keywords"`
}

type Keyword struct {
	Keyword  string
	Category string
}

func init() {
	if data, err := f.ReadFile("config.yaml"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		if err = yaml.Unmarshal(data, &Conf); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
