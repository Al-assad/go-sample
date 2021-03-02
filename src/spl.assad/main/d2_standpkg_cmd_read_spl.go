package main

import (
	"bufio"
	"fmt"
	"os"
)

// 从控制台读取输入（os.Stdin）

func main() {
	//usePrintln()
	readStdin()

}

// 使用 fmt.Scanln 从标准输入读取一行， 使用空格分隔
func usePrintln() {
	var name string
	var city string

	fmt.Println("input your name: ")
	fmt.Scanln(&name)
	fmt.Println("hello,", name)

	fmt.Println("input your name and city")
	fmt.Scanln(&name, &city)
	fmt.Println("hello,", name, "in", city)

	//input your name:
	//assad
	//hello, assad
	//input your name and city
	//assad guangzhou
	//hello, assad in guangzhou
}

// 从标准输入 os.Stdin 读取
func readStdin() {
	fmt.Println("input something: ")

	// 读取标准输入
	var inputReader *bufio.Reader = bufio.NewReader(os.Stdin)
	// 从标准输入中读取字符串
	input, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Println("echo > ", input)
	}
}
