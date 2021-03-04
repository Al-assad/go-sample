package main

import (
	"fmt"
	"os"
	"runtime"
)

// @desc go 显式调用 runtime gc

func main() {
	inFile, _ := os.Open(".gitignore")

	// 获取 runtime 当前内存状态
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("%v Kb \n", mem.Alloc/1024)

	// 设置 inFile 内存回收前的操作
	runtime.SetFinalizer(inFile, func(f *os.File) {
		f.Close()
	})

	// 显式触发 GC
	runtime.GC()

}
