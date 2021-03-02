package main

import (
	"fmt"
	"reflect"
)

// interface 接口使用示例

type Biology interface {
	Eat() string
}

type Animal interface {
	Move() string
}

type Cat struct {
}

// Cat 实现 Biology.Eat() 接口
func (c *Cat) Eat() string {
	return "fish"
}

// Cat 实现 Animal.Move() 接口
func (c *Cat) Move() string {
	return "slunk ~~"
}

func main() {
	cat := new(Cat)
	fmt.Println(cat.Eat())
	fmt.Println(cat.Move())

	// 使用接口指向接口体
	var bio Biology
	bio = cat
	fmt.Println(bio.Eat())
	var ani Animal
	ani = cat
	fmt.Println(ani.Move())

	// 通过接口声明实现类
	var bio2 Biology = new(Cat)
	fmt.Println(bio2.Eat())

	// 接口作为函数入参
	descBiology(bio2)
}

// 接口作为函数入参数
func descBiology(bio Biology) {
	fmt.Printf("What %s Eatting？%s! \n", reflect.TypeOf(bio).String(), bio.Eat())
}
