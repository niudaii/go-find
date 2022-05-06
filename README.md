## go-find
花了半天时间写的一个的文件名、文件内容搜索工具，主要是为了hvv中快速发现敏感文件、配置文件等。
## Usage

用法

```
./go-find_darwin_amd64 -h
./go-find_darwin_amd64 find -h	(文件名搜索)
./go-find_darwin_amd64 findstr -h	（文件内容搜索）
./go-find_darwin_amd64 find -d ./test	（指定路径下所有文件名查找关键词）
./go-find_darwin_amd64 findstr -d ./test	（指定路径下所有文件内容查找关键词）
```

全部参数

```
	-d, --dir string               dir to search (default "./")
      --file-white-ext string    file white ext, split by comma （文件名后缀白名单）
  -h, --help                     help for go-find
  -k, --keywords string          keywords, split by comma
  -o, --output string            file to write output[xxx.txt] (default "gofind_result.txt")
      --path-black-word string   path black word, split by comma （路径黑名单）
```

## 使用截图

![image-20220506135149384](https://nnotes.oss-cn-hangzhou.aliyuncs.com/notes/image-20220506135149384.png)

测试目录

![image-20220506135245769](https://nnotes.oss-cn-hangzhou.aliyuncs.com/notes/image-20220506135245769.png)

搜索文件名

![image-20220506135327457](https://nnotes.oss-cn-hangzhou.aliyuncs.com/notes/image-20220506135327457.png)

搜索文件内容

![image-20220506135355486](https://nnotes.oss-cn-hangzhou.aliyuncs.com/notes/image-20220506135355486.png)

`--file-white-ext` 增加文件名后缀白名单

![image-20220506135447332](https://nnotes.oss-cn-hangzhou.aliyuncs.com/notes/image-20220506135447332.png)

`--path-black-word` 增加路径黑名单

![image-20220506135523395](https://nnotes.oss-cn-hangzhou.aliyuncs.com/notes/image-20220506135523395.png)

## 说明

- 内置了文件名关键词、文件内容关键词、路径黑名单、文件名后缀白名单，用户指定 keyword 参数时不加载默认关键词。
- 文件内容搜索时，匹配到时打印出前后相邻的50个字符。
- `config/config.yaml` 中可自行配置各种关键词。
- 可以使用 `|` 分隔关键词和分类，比如 `weaver.properties|泛微`，类型会在匹配到的时候打印出来，因此关键词中不要出现 `|`
