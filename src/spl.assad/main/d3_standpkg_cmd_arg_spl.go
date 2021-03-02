package main

import (
	"fmt"
	"os"
	"strings"
)

// @desc 从启动命令行读取参数 os.Args

func main() {
	//  获取参数列表
	if len(os.Args) > 1 {
		fmt.Println("args: ", strings.Join(os.Args[1:], " "))
	} else {
		fmt.Println("app", os.Args[0])
	}
}
