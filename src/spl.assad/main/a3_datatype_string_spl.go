package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

// @desc 字符串类型使用

func main() {
	print("hello world \n")
	println(`hello world \n`) // 原样输出
	//runeCount()
	//string_opera()
	//coverNum()
	//charAtIndex()
	//changeCharAtIndex()
	//regexTest()
}

// 统计字符长度
func runeCount() {
	str1 := "Are you ok ?"
	println(len(str1))                    //12
	println(utf8.RuneCountInString(str1)) //12

	str2 := "你还好吗？"
	println(len(str2))                    //15
	println(utf8.RuneCountInString(str2)) //5
}

// 字符串 <-> 数值 转换
func coverNum() {
	// 数值转换 string
	var num int = 233
	str := strconv.Itoa(num)
	println(str)

	var num2 float64 = 233.33
	str = strconv.FormatFloat(num2, 'f', 4, 64)
	println(str)

	// string 转换 num
	num3, err := strconv.Atoi("23")
	println(num3, err)
	num4, _ := strconv.ParseInt("233", 10, 0)
	println(num4)

	num5, _ := strconv.ParseFloat("23.33", 64)
	println(num5)
}

// 字符串常用操作
func string_opera() {
	str := "Let aeroplanes circle moaning overhead, Scribbling on the sky the message He Is Dead."

	// 前后缀判断
	r := strings.HasPrefix(str, "Le")
	println(r)
	r = strings.HasSuffix(str, "Dead.")
	println(r)

	// 字符串包含
	r = strings.Contains(str, "mess")
	println(r)
	r = strings.ContainsAny(str, "a")
	println(r)

	// 获取子串或字符出现位置
	r2 := strings.Index(str, "over")
	println(r2)

	// 字符串替换
	r3 := strings.ReplaceAll(str, "He", "She")
	println(r3)

	// 统计字符串出现次数
	r4 := strings.Count(str, "a")
	println(r4)

	// 生成重复字符串
	r5 := strings.Repeat("ha", 20)
	println(r5)

	// 修改字符串大小写
	r6 := strings.ToUpper(str)
	println(r6)
	r6 = strings.ToLower(str)
	println(r6)

	// 裁剪字符串
	r7 := strings.TrimSpace(" are you ok? ")
	println(r7)
	r7 = strings.Trim("> are you ok? >>", ">")
	println(r7)

	// 分割字符串
	r8 := strings.Split(str, " ")
	for row := range r8 {
		print(r8[row] + "-")
	}
	println()

	// 拼接 slice 到字符串
	r9 := strings.Join(r8, " ")
	println(r9)

}

// 获取字符串指定下标字符
func charAtIndex() {
	// ascii 字符串
	var str = "hello world"
	fmt.Println(string(str[0])) // a
	for i := 0; i < len(str); i++ {
		fmt.Println(string(str[i]))
	}

	// utf8 字符串，使用 rune 数组
	str2 := "hello 世界"
	strRune := []rune(str2)
	fmt.Println(string(strRune[6])) // 世
	for i := 0; i < len(strRune); i++ {
		fmt.Println(string(strRune[i]))
	}

}

// 修改字符串指定下标字符
func changeCharAtIndex() {
	// ascii 字符串
	str := "hello"
	chs := []byte(str)
	chs[2] = 'A'
	strChange := string(chs)
	fmt.Println(strChange) // heAlo

	// utf8 字符串，使用 rune 数组
	str2 := "hello 世界"
	runes := []rune(str2)
	runes[6] = '视'
	str2Change := string(runes)
	fmt.Println(str2Change) // hello 视界
}

// 正则匹配
func regexTest() {
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
}
