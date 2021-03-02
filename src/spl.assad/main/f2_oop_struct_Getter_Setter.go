package main

import "fmt"

// go struct 结构体示例（字段 Getter、Setter）

// 简单结构体
type User3 struct {
	name  string
	scope int
}

func (u *User3) GetName() string {
	return u.name
}

func (u *User3) GetScope() int {
	return u.scope
}

func (u *User3) SetName(name string) {
	u.name = name
}

func (u *User3) setScope(scope int) {
	u.scope = scope
}

func main() {
	u := &User3{"Assad", 1024}
	fmt.Println(u.GetName(), u.GetScope())
	u.SetName("Alex")
	u.setScope(2048)
	fmt.Println(u.GetName(), u.GetScope())
	//Assad 1024
	//Alex 2048
}
