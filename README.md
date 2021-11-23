## go-find
花了半天时间写的一个的文件名、文件内容搜索工具，主要是为了hvv中快速发现敏感文件、配置文件等。
## Usage

用法

```
./go-find_darwin_amd64 -h
./go-find_darwin_amd64 find -h		(文件名搜索)
./go-find_darwin_amd64 findstr -h	（文件内容搜索）
./go-find_darwin_amd64 find -d /Users/niudai/goland/src/gin-osint	（指定路径下所有文件名查找关键词）
./go-find_darwin_amd64 findstr -d /Users/niudai/goland/src/gin-osint	（指定路径下所有文件内容查找关键词）
```

全部参数

```
-o, --output string     output name (default "matchResults.txt")
-d, --dir string        dir to search
    --fwe string        filenameWhiteExt, split by comma
    --pbw string        pathBlackWord, split by comma
```

## 使用截图

搜索文件名

![image-20211123151832477](https://nnotes.oss-cn-hangzhou.aliyuncs.com/notes/image-20211123151832477.png)

搜索文件内容

![image-20211123151844470](https://nnotes.oss-cn-hangzhou.aliyuncs.com/notes/image-20211123151844470.png)

`--pbw`增加路径黑名单

![image-20211123151334211](https://nnotes.oss-cn-hangzhou.aliyuncs.com/notes/image-20211123151334211.png)

`--fwe`增加文件名后缀白名单

![image-20211123151523871](https://nnotes.oss-cn-hangzhou.aliyuncs.com/notes/image-20211123151523871.png)

## 说明

- 内置了文件名关键词、文件内容关键词、路径黑名单、文件名后缀白名单，用户输入的话会进行去重合并。
- 文件内容搜索时，匹配到时只打印出前后相邻的10个字符。
- `config/config.yaml`中可自行配置各种关键词。

