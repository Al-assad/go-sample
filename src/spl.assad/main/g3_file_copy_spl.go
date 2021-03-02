package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

// 文件拷贝示例

func main() {
	//copy1()
	//copy2()
	copy3()
}

// 文件拷贝 - 使用 Reader、Writer 进行逐行拷贝
func copy1() {
	inFile, _ := os.OpenFile("file_test/poem.dat", os.O_RDONLY, 0)
	outFile, _ := os.OpenFile("file_test/poem_copy.dat", os.O_CREATE|os.O_WRONLY, 0666)
	defer inFile.Close()
	defer outFile.Close()

	reader := bufio.NewReader(inFile)
	writer := bufio.NewWriter(outFile)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		writer.WriteString(line)
	}
	writer.Flush()
}

// 文件拷贝 - 使用 Reader、Writer 进行逐行字节拷贝
func copy2() {
	inFile, _ := os.OpenFile("file_test/poem.dat", os.O_RDONLY, 0)
	outFile, _ := os.OpenFile("file_test/poem_copy2.dat", os.O_CREATE|os.O_WRONLY, 0666)
	defer inFile.Close()
	defer outFile.Close()

	reader := bufio.NewReader(inFile)
	writer := bufio.NewWriter(outFile)

	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		writer.Write(line)
	}
	writer.Flush()
}

// 文件拷贝 - 使用 ioutil 进行小文件快速拷贝（全部读取、全部写入）
func copy3() {
	inData, _ := ioutil.ReadFile("file_test/poem.dat")
	ioutil.WriteFile("file_test/poem_copy3.dat", inData, 0666)
}
