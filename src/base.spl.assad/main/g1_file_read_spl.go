package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// @desc 文件读取示例

func main() {
	readFile()
	//readAllFile()
	//readAllReader()
	//readFileWithBuff()
}

// 文件读取
func readFile() {
	// 打开文件
	file, err := os.Open(".gitignore")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 关闭文件
	defer file.Close()

	// 使用 Reader 读取文件
	reader := bufio.NewReader(file)
	// 循环读取文本
	for {
		line, err := reader.ReadString('\n')
		fmt.Print(line)
		if err == io.EOF {
			break
		}
	}
}

// ioutil 读取整个文件
func readAllFile() {
	inFilePath := ".gitignore"
	// 读取整个文件
	buf, err := ioutil.ReadFile(inFilePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(buf))
}

// ioutil 读取整个 Reader
func readAllReader() {
	file, _ := os.Open(".gitignore")
	reader := bufio.NewReader(file)
	// 读取整个 Reader
	buf, _ := ioutil.ReadAll(reader)
	fmt.Println(string(buf))
}

// 自定义缓冲块读取文件
func readFileWithBuff() {
	file, _ := os.Open(".gitignore")
	defer file.Close()

	count := 0
	// 定义一块 1kb 的缓存块
	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf) //n 表示读到的字节数
		count++
		if n == 0 {
			break
		}
	}
	fmt.Println("Read 1kb buffer count:", count)
}
