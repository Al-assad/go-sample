package main

import (
	myfoo "base.spl.assad/func_foo"
	"base.spl.assad/main/mypack"
	"fmt"
	"runtime"
)

// @desc 函数调用示例

func main() {
	//helloWorld()
	//test1()
	//test2()
	//test3()
	testImHere()
}

// 文件内函数
func helloWorld() {
	fmt.Println("hello world!")
	fmt.Println(runtime.Version())
}

// 包内函数 a_foo.go 调用
func test1() {
	var r1 = sum(1, 2)
	var r2 = Sum(3, 4)
	fmt.Println(r1 + r2)
}

// 包外函数 func_foo/foo.go 调用
func test2() {
	var r = myfoo.Minus(10, 20)
	fmt.Println(r)
}

// 调用子包函数
func testImHere() {
	mypack.ImHere()
}
