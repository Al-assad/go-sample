package main

import (
	"fmt"
	"runtime"
	myfoo "site.nobody/func_foo"
	"spl.assad/func_foo"
	"spl.assad/main/mypack"
)

// 函数调用测试

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

// 包内函数 foo.go 调用
//func test1() {
//	var r1 = sum(1, 2)
//	var r2 = Sum(3, 4)
//	fmt.Println(r1 + r2)
//}

// 包外函数 func_foo/foo.go 调用
func test2() {
	var r = func_foo.Minus(10, 20)
	fmt.Println(r)
}

// 包外函数重命名调用
func test3() {
	var r = myfoo.Minus(10, 20)
	fmt.Println(r)
}

// 调用子包函数
func testImHere() {
	mypack.ImHere()
}
