package main

import (
	"fmt"
	"regexp"
	"strings"
)

// @desc go regexp 正则匹配

func main() {
	str := "John: 2578.34 William: 4567.23 Steve: 5632.18" // 目标字符串
	pattern := "[0-9]+.[0-9]+"                             // 正则表达式

	// 快速判断字符串正则匹配
	ok1, _ := regexp.Match(pattern, []byte(str))
	ok2, _ := regexp.MatchString(pattern, str)
	fmt.Println(ok1, ok2) // true true

	// 正则查找所有匹配项
	re, _ := regexp.Compile(pattern)
	maStrs := re.FindAllString(str, len(str))
	for _, maStr := range maStrs {
		fmt.Print(maStr, " ")
	}
	println() // 2578.34 4567.23 5632.18

	// 正则替换匹配项
	r1 := re.ReplaceAllString(str, "####")
	fmt.Println(r1) // John: #### William: #### Steve: ####

	// 正则替换匹配项，使用函数指定替换行为
	replaceFunc := func(s string) string {
		return strings.ReplaceAll(s, ".", "#")
	}
	r2 := re.ReplaceAllStringFunc(str, replaceFunc)
	fmt.Println(r2) // John: 2578#34 William: 4567#23 Steve: 5632#18

	// 查找子匹配项
	str2 := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pattern2 := "Steve:\\s*([0-9]+.[0-9]+)\\s*"
	re2, _ := regexp.Compile(pattern2)
	ms := re2.FindStringSubmatch(str2)
	fmt.Println(ms[1]) // 5632.18

}
