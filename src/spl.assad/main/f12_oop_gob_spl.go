package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// @desc go gob 二进制序列化、反序列化
// Gob 是 Go 自己的以二进制形式序列化和反序列化程序数据的格式, G
// Gob 特定地用于纯 Go 的环境中，例如，两个用 Go 写的服务之间的通信。这样的话服务可以被实现得更加高效和优化

// 发送方结构体
type T struct {
	X int
	Y int
	Z int
}

// 获取方结构体
type U struct {
	X int
	Y int
}

func main() {
	var network bytes.Buffer
	// gob 编码器
	enc := gob.NewEncoder(&network)
	// gob 解码器
	dec := gob.NewDecoder(&network)

	// 序列化编码
	t := &T{1, 2, 3}
	enc.Encode(t)
	// 反序列化解码
	var u U
	dec.Decode(&u)

	fmt.Println("T", *t) // T {1 2 3}
	fmt.Println("U", u)  // U {1 2}

}
