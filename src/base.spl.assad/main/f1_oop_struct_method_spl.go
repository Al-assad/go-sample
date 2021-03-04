package main

import (
	"base.spl.assad/func_foo"
	"fmt"
	"strconv"
)

// @desc go struct 结构体示例

// 简单结构体
type User struct {
	name  string
	city  string
	scope int
}

// User 结构体工厂方法
func NewUser(name string, city string) *User {
	u := new(User)
	u.name = name
	u.city = city
	u.scope = 1000
	return u
}

// User 结构体 String 方法
func (us User) String() string {
	return "User{name=" + us.name + ", city=" + us.city + ", scope=" + strconv.Itoa(us.scope) + "}"
}

// User 自定义方法
func (us *User) DownScope(val int) {
	us.scope = us.scope - val
}

func main() {
	//structDefine()
	//structDirectInit()
	//structFuncParamTest()
	structMethodTest()
}

// 声明、赋值、访问结构体
func structDefine() {
	// 声明结构体，此处 u 为 struct 指针
	u := new(User)
	// 结构体字段赋值
	u.name = "assad"
	u.city = "guangzhou"
	u.scope = 233
	// 访问结构体字段
	fmt.Println(u.name, u.city, u.scope) // assad guangzhou 233
	// 打印结构体所有字段内容
	fmt.Printf("%v \n", *u) // {assad guangzhou 233}

	// 声明其他包的结构体
	m := new(func_foo.Man)
	m.Name = "Alex"
	// m.scope 包外不可访问
	fmt.Println(m.Name)
	fmt.Printf("%v \n", *m) // {Alex 0}
}

// 结构体声明后直接赋值
func structDirectInit() {
	// 此处 u1 为一个 struct 指针
	u1 := &User{"assad", "guangzhou", 1024}
	fmt.Printf("%v \n", *u1) // {assad guangzhou 1024}

	// 此处 u2 为一个 struct 实例
	var u2 User
	u2 = User{"alex", "shanghai", 2331}
	// 等同于
	u3 := User{"alex", "shanghai", 2331}
	fmt.Printf("%v \n", u2) // {alex shanghai 2331}
	fmt.Printf("%v \n", u3)
}

// 结构作为方法入参数、出参
func structFuncParamTest() {
	u1 := &User{"assad", "guangzhou", 1024}
	u2 := upUserScope(u1)
	fmt.Printf("%v \n", *u2) // {assad guangzhou 1124}
	fmt.Printf("%v \n", *u1) // {assad guangzhou 1124}
}
func upUserScope(u *User) (user *User) {
	u.scope = u.scope + 100
	return u
}

// 结构体工厂方法、String 方法、自定义方法使用
func structMethodTest() {
	u := NewUser("Bob", "Beijing")
	fmt.Println(u.scope) // 1000
	u.DownScope(100)
	fmt.Println(u.scope) // 900
}
