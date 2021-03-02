package main

import (
	"fmt"
	"reflect"
)

// go 类型反射示例（TypeOf，ValueOf）

type User5 struct {
	name  string
	city  string
	Scope int `lim:"102400" desc:"scope of top"` // tag 中规定 key:value 分组
}

func (u User5) Say() {
	fmt.Println("hello world!")
}

func (u User5) SayWhat(what string) string {
	return "Hello, " + what
}

func main() {
	basicTypeReflect()
	structReflect()
	methodCall()
	funcCall()
}

func pickTwo(a, b int) int {
	return a + b
}

// 基本类型反射
func basicTypeReflect() {
	var num float64 = 233.33

	// 获取基本类型的 Type 反射
	var t reflect.Type = reflect.TypeOf(num)
	fmt.Println(t.Name())   // float64
	fmt.Println(t.String()) // float64

	// 获取基本类型的 Value 反射
	var v reflect.Value = reflect.ValueOf(num)
	fmt.Println(v.Type().Name())   // float64
	fmt.Println(v.String())        // <float64 Value>
	fmt.Println(v.IsZero())        // false
	fmt.Println(v.Kind().String()) // float64
}

// 结构体反射
func structReflect() {
	// 结构体反射
	user := User5{"assad", "guangzhou", 233}
	ut := reflect.TypeOf(user)
	uv := reflect.ValueOf(user)
	fmt.Println(ut.Name())          // User5
	fmt.Println(ut.String())        //main.User5
	fmt.Println(ut.Kind().String()) // struct
	fmt.Println(uv.IsZero())        // false
	// 遍历属性
	for i := 0; i < ut.NumField(); i++ {
		f := ut.Field(i)
		fmt.Print(f.Name, " ") // name city Scope
	}
	println()

	// 按照属性名获取属性值
	f := uv.FieldByName("city")
	fmt.Println(f.String()) // guangzhou

	// 按照方法名获取方法
	mt, _ := ut.MethodByName("Say")
	mv := uv.MethodByName("Say")
	fmt.Println(mt.Name)                   // Say
	fmt.Println(mt.Type.String())          // func(main.User5) string
	fmt.Println(mv.Kind() == reflect.Func) // true
}

// 方法反射调用
func methodCall() {
	user := User5{"assad", "guangzhou", 233}
	userType := reflect.ValueOf(user)

	// 调用方法 - 无入参
	methodSay := userType.MethodByName("Say")
	methodSay.Call([]reflect.Value{}) // hello world!

	// 调用方法 - 有入参数
	methodSayWhat := userType.MethodByName("SayWhat")
	params := []reflect.Value{
		reflect.ValueOf("World"),
	}
	rs := methodSayWhat.Call(params)
	r := rs[0].String()
	fmt.Println(r) // Hello, World
}

func funcCall() {
	myFunc := reflect.ValueOf(pickTwo)
	params := []reflect.Value{
		reflect.ValueOf(30),
		reflect.ValueOf(50),
	}
	rs := myFunc.Call(params)
	fmt.Println(rs[0].Int()) // 80
}
