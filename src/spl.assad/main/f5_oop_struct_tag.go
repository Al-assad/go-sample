package main

import (
	"fmt"
	"reflect"
)

// go struct 标签示例

type User4 struct {
	name  string "this is username"
	city  string "city of user"
	scope int    `lim:"102400" desc:"scope of top"` // tag 中规定 key:value 分组
}

func main() {
	u := User4{"assad", "guangzhou", 233}
	// 获取 u 所有字段 tag
	uType := reflect.TypeOf(u)
	for i := 0; i < uType.NumField(); i++ {
		uf := uType.Field(i)
		fmt.Printf("%v \n", uf.Tag)
	}
	//this is username
	//city of user
	//lim:"102400" desc:"scope of top"

	// 获取 u.scope 指定 tag key 信息
	lim := uType.Field(2).Tag.Get("lim")
	fmt.Println(lim) // 102400

}
