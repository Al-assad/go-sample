package main

import "fmt"

// go interface 接口嵌套示例

type In interface {
	Set(param int)
}

type Out interface {
	Get() int
}

// Operator 接口包含 In Out 嵌套接口
type Operator interface {
	Handle()
	In
	Out
}

type IncreOperator struct {
	val int
}

// 实现接口方法
func (o *IncreOperator) Handle() {
	o.val = o.val + 1
}
func (o *IncreOperator) Get() int {
	return o.val
}
func (o *IncreOperator) Set(param int) {
	o.val = param
}

func main() {
	o := new(IncreOperator)

	var oper Operator = o
	oper.Set(20)
	oper.Handle()
	r := oper.Get()

	fmt.Println(r)
}
