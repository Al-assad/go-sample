package main

import (
	"fmt"
	"github.com/usk81/generic"
)

// 第三方包调用示例

// 获取第三方库 https://github.com/usk81/generic
// go get github.com/usk81/generic
func main() {
	v := 1.0
	var tb generic.Bool
	tb.Set(v)
	vb := tb.Weak()
	fmt.Printf("%v, (%T)\n", vb, vb)

	var tf generic.Float
	tf.Set(v)
	vf := tf.Weak()
	fmt.Printf("%v, (%T)\n", vf, vf)

	var ti generic.Int
	ti.Set(v)
	vi := ti.Weak()
	fmt.Printf("%v, (%T)\n", vi, vi)

}
