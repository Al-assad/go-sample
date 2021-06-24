package main

import "fmt"

// @desc 空 interface 使用示例

// 如果定义一个没有任何方法的空接口，该接口可以接收任何类型参数
func main() {

	// slice 使用空元素
	sl := make([]interface{}, 5)
	sl[0] = 1
	sl[1] = 233.2
	sl[2] = true
	sl[3] = "hello"
	for _, e := range sl {
		fmt.Println(e)
	}

	// map 使用空元素
	ma := make(map[string]interface{})
	ma["name"] = "assad"
	ma["score"] = 233
	ma["isdev"] = true
	for k, v := range ma {
		fmt.Println(k, "=", v)
	}

}
