package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 文件写入示例

func main() {
	writeFile()
	//writeAllString()
}

// 文件写入
func writeFile() {
	// 打开文件，只允许写入、不存在即创建
	// os.O_RDONLY：只读
	// os.O_WRONLY：只写
	// os.O_RDWR：读写
	// os.O_CREATE：创建：如果指定文件不存在，就创建该文件
	// os.O_TRUNC：截断：如果指定文件已存在，就将该文件的长度截为 0
	// os.O_APPEND: 追加
	file, err := os.OpenFile("file_test/poem.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()
	// 装载入 writer
	writer := bufio.NewWriter(file)
	// 写入 string
	writer.WriteString("Whenever you need me， I'll be here\n")
	writer.WriteString("Whenever you're in trouble， I'm always near\n")
	writer.WriteString("Whenever you feel alone， and you think everyone has given up\n")

	// 输出 writer 缓冲到磁盘
	writer.Flush()
}

// ioutil 写入整个字符串
func writeAllString() {
	poem := "I have had my invitation to this world's festival, " +
		"and thus my life has been blessed. \n" +
		"Early in the day it was whispered that we should sail in a boat, \n" +
		"only thou and I, and never a soul in the world would know of this our pilgrimage to no country and to no end.\n"

	outFilePath := "file_test/poem2.dat"
	// ioutil.WriteFile 文件语义等同于 os.O_CREATE ｜ os.O_TRUNC
	ioutil.WriteFile(outFilePath, []byte(poem), 0666)
}
