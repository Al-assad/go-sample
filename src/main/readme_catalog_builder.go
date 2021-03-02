package main

// README 目录构建工具

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	fileNoteRe     *regexp.Regexp // go 文件注释匹配正则表达式
	specFileNameRe *regexp.Regexp // 特殊 go 文件名匹配正则表达式
	pkgDicts       []PkgDesc      // 包描述字典列表
)

func init() {
	fileNoteRe, _ = regexp.Compile("//.*@desc (.+)")
	specFileNameRe, _ = regexp.Compile("([a-zA-Z])+([0-9]+)_(.*)")
	pkgDicts = []PkgDesc{
		{"./src/spl.assad/main",
			"Go 基本使用示例代码",
			map[string]string{
				"a": "Go 数据类型、数据结构",
				"b": "Go 控制结构语法",
				"c": "Go 函数",
				"d": "Go 标准包",
				"f": "Go 面向对象",
				"g": "Go 文件处理",
				"h": "Go 异常处理",
			},
		},
		// other pkg desc
	}
}

// go 文件结构
type GoItem struct {
	Name  string
	Path  string
	Note  string
	Group string
}

// 包描述结构
type PkgDesc struct {
	PkgPath   string
	Desc      string
	GroupDesc map[string]string
}

func main() {
	BuildAllCatalog()
}

// 构建字典中所有的包 README 索引，并输出到终端
func BuildAllCatalog() {
	for _, pkg := range pkgDicts {
		catalog := BuildCatalog(pkg)
		fmt.Println(catalog)
	}
}

// 构建指定包 README 目录索引
func BuildCatalog(pkgDesc PkgDesc) string {
	targetPath := pkgDesc.PkgPath
	resultContent := buildPkgMdLine(pkgDesc.Desc)

	// 获取目录下所有文件
	dir, err := os.Open(targetPath)
	if err != nil {
		fmt.Println("no exists ", targetPath)
		fmt.Println(err.Error())
		return ""
	}
	fileInfos, _ := dir.Readdir(0)

	// 遍历文件，提取 GoItems 要素
	goItems := make([]GoItem, 0)
	for _, fileInfo := range fileInfos {
		// 目录不处理
		if fileInfo.IsDir() {
			continue
		}
		// go 文件路径
		filePath := targetPath + "/" + fileInfo.Name()
		file, _ := os.Open(filePath)
		note := getFirstNoteFromGoFile(file)
		if note == "" {
			continue
		}
		// 组装 markdown 要素
		goItem := GoItem{fileInfo.Name(), filePath, note, ""}
		goItems = append(goItems, goItem)
	}

	// 按照文件排序
	sort.Slice(goItems, func(i, j int) bool {
		m := specFileNameRe.FindStringSubmatch(goItems[i].Name)
		n := specFileNameRe.FindStringSubmatch(goItems[j].Name)
		if len(m) >= 3 && len(n) >= 3 {
			if m[1] < n[1] {
				return true
			} else if m[1] > n[1] {
				return false
			} else {
				im2, _ := strconv.Atoi(m[2])
				in2, _ := strconv.Atoi(n[2])
				return im2 < in2
			}
		} else {
			return goItems[i].Name < goItems[j].Name
		}
	})

	// 文件名格式化
	for i, item := range goItems {
		fname, group := formatFileName(item.Name)
		goItems[i].Name = fname
		goItems[i].Group = group
	}

	// 组装 Markdown 文本
	var mdlines []string
	if pkgDesc.GroupDesc != nil && len(pkgDesc.GroupDesc) > 0 {
		mdlines = buildMdLinesWithGroup(&goItems, pkgDesc.GroupDesc)
	} else {
		mdlines = buildMdLines(&goItems)
	}

	resultContent = resultContent + "\n" + strings.Join(mdlines, "\n") + "\n<br>"
	return resultContent
}

// 获取 go 文件头注释
func getFirstNoteFromGoFile(file *os.File) string {
	reader := bufio.NewReader(file)
	for {
		lineBytes, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		notes := fileNoteRe.FindStringSubmatch(lineBytes)
		if len(notes) > 0 {
			return strings.TrimSpace(notes[1])
		}
	}
	return ""
}

// 格式化文件名称
func formatFileName(name string) (fName string, group string) {
	match := specFileNameRe.FindStringSubmatch(name)
	if len(match) > 0 {
		return match[3], match[1]
	} else {
		return name, ""
	}
}

// GoItem 列表转化为 MD 文本行
func buildMdLines(goItems *[]GoItem) []string {
	mdLines := make([]string, 0)
	for _, item := range *goItems {
		line := "* [" + item.Name + "](" + item.Path + "): " + item.Note
		mdLines = append(mdLines, line)
	}
	return mdLines
}

// GoItem 列表转化为 MD 文本行，增加分组行描述
func buildMdLinesWithGroup(goItems *[]GoItem, groupDict map[string]string) []string {
	curGroup := ""
	mdLines := make([]string, 0)
	for _, item := range *goItems {
		// 插入分组描述行
		if curGroup != item.Group {
			groupDesc := groupDict[item.Group]
			mdLines = append(mdLines, buildPkgGroupMdLine(groupDesc))
			curGroup = item.Group
		}
		line := "* [" + item.Name + "](" + item.Path + "): " + item.Note
		mdLines = append(mdLines, line)
	}
	return mdLines

}

// 转化包描述为 md 行
func buildPkgMdLine(desc string) string {
	return "### " + desc
}

// 转化包分组描述为 md 行
func buildPkgGroupMdLine(desc string) string {
	return "#### " + desc
}
