package main

import (
	"encoding/xml"
	"fmt"
)

// @desc go xml 处理

// https://golang.org/pkg/encoding/xml/

type Person2 struct {
	Name     string
	Gender   int
	IsNewMan bool
}

// 使用 tag 规定映射关系
type MyUser struct {
	FirstName string   `xml:"Name>First"`
	LastName  string   `xml:"Name>Last"`
	Email     []string `xml:"Email>Value"`
}

func main() {
	baseSample()
	useTagSample()
}

func baseSample() {
	// xml 序列化
	person := &Person2{"assad", 0, true}

	// 直接序列化
	xs, _ := xml.Marshal(person)
	fmt.Println(string(xs))
	// <Person2><Name>assad</Name><Gender>0</Gender><IsNewMan>true</IsNewMan></Person2>

	// 序列化并使缩进格式
	xs2, _ := xml.MarshalIndent(person, "", "  ")
	fmt.Println(string(xs2))
	// <Person2>
	//   <Name>assad</Name>
	//   <Gender>0</Gender>
	//   <IsNewMan>true</IsNewMan>
	// </Person2>

	// xml 反序列化
	var person2 Person2
	xml.Unmarshal(xs, &person2)
	fmt.Println(person2)
	// {assad 0 true}
}

// 使用 struct 描述 xml 层级关系
func useTagSample() {
	user := &MyUser{"Assad", "Al", []string{"areyouok@sample.com", "darkfantasy@sample.con"}}

	// 序列化
	xs, _ := xml.MarshalIndent(user, "", "  ")
	fmt.Println(string(xs))
	// <MyUser>
	//   <Name>
	//     <First>Assad</First>
	//     <Last>Al</Last>
	//   </Name>
	//   <Email>
	//     <Value>areyouok@sample.com</Value>
	//     <Value>darkfantasy@sample.con</Value>
	//   </Email>
	// </MyUser>

	// 反序列化
	var user2 MyUser
	xml.Unmarshal(xs, &user2)
	fmt.Println(user2)
	// {Assad Al [areyouok@sample.com darkfantasy@sample.con]}
}
