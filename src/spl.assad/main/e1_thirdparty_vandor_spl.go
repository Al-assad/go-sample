package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
)

// 第三方包调用示例

// 使用 https://github.com/deckarep/golang-set 实现的 set 结构

// 第三方依赖安装方式：
// 方式 1：go get
//        GOPATH=${project_path} && cd ${project_path}/src
//        go get github.com/deckarep/golang-set

// 方式 2：vendor
//        GOPATH=${project_path} && cd ${project_path}/src
//		  govendor init
//        govendor fetch github.com/deckarep/golang-set

//        已存在 ./vendor 目录，根据 .vendor/vendor.json 同步所有第三方库元源码
//        govendor sync

func main() {
	mlist := []int{1, 2, 3, 4, 5, 6, 7, 7, 2, 1, 2}
	mset := mapset.NewSet()
	for _, e := range mlist {
		mset.Add(e)
	}
	fmt.Println("size:", mset.Cardinality())
	fmt.Println(mset.Contains(3))
	fmt.Println(mset.String())
	for e := range mset.Iter() {
		fmt.Print(e, " ")
	}
	println()

}
