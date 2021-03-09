package main

// @desc jsonparser 高性能 json 库使用（灵活结构操作）
// jsonparser 适用于灵活 json 结构的解析组装
import (
	"fmt"
	"github.com/Al-assad/go-sample/src/tool.spl.assad/model"
	"github.com/buger/jsonparser"
	"strconv"
)

var jsonData = []byte(`{
  "person": {
    "name": {
      "first": "Leonid",
      "last": "Bugaev",
      "fullName": "Leonid Bugaev"
    },
    "github": {
      "handle": "buger",
      "followers": 109
    },
    "avatars": [
      { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" }
    ]
  },
  "company": {
    "name": "Acme"
  }
}`)

func main() {
	getJsonField()
	modifyJsonField()
}

// 自由获取 json 字段
func getJsonField() {
	// 获取 person.name.first
	firstName, err := jsonparser.GetString(jsonData, "person", "name", "first")
	model.CheckErr(err)
	fmt.Println(firstName) // Leonid

	// 获取 person.github.follows
	followers, err := jsonparser.GetInt(jsonData, "person", "github", "followers")
	model.CheckErr(err)
	fmt.Println(followers) // 109

	// 获取 person.avatars[0].url
	url, err := jsonparser.GetString(jsonData, "person", "avatars", "[0]", "url")
	model.CheckErr(err)
	fmt.Println(url) // https://avatars1.githubusercontent.com/u/14009?v=3&s=460

	// arrayeach 获取 person.avators 所有元素
	_, err = jsonparser.ArrayEach(jsonData, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		url, _ := jsonparser.GetString(value, "url")
		utype, _ := jsonparser.GetString(value, "type")
		fmt.Println(url, utype)
	}, "person", "avatars")
	model.CheckErr(err)
	// https://avatars1.githubusercontent.com/u/14009?v=3&s=460 thumbnail

	// objecteach 获取 person.name 所有元素
	err = jsonparser.ObjectEach(jsonData, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		fmt.Printf("%s = %s\n", string(key), string(value))
		return nil
	}, "person", "name")
	/*
		first = Leonid
		last = Bugaev
		fullName = Leonid Bugaev
	*/

	// 批量获取 Key
	paths := [][]string{
		[]string{"person", "name", "fullName"},
		[]string{"person", "avatars", "[0]", "url"},
		[]string{"company", "name"},
	}
	jsonparser.EachKey(jsonData, func(i int, bytes []byte, valueType jsonparser.ValueType, err error) {
		switch i {
		case 0:
			fmt.Println("person.name.fullName = ", string(bytes))
		case 1:
			fmt.Println("person.avatars.[0].url = ", string(bytes))
		case 2:
			fmt.Println("company.name = ", string(bytes))
		}
	}, paths...)
	/*
		person.name.fullName =  Leonid Bugaev
		person.avatars.[0].url =  https://avatars1.githubusercontent.com/u/14009?v=3&s=460
		company.name =  Acme
	*/
}

// 自由修改 json 属性
func modifyJsonField() {

	// 修改 person.github.followers
	followers, err := jsonparser.GetInt(jsonData, "person", "github", "followers")
	model.CheckErr(err)
	followers += 100
	rjson, err := jsonparser.Set(jsonData, []byte(strconv.Itoa(int(followers))), "person", "github", "followers")
	model.CheckErr(err)
	fmt.Println(string(rjson))

	// 删除 person.github.followers
	rjson = jsonparser.Delete(jsonData, "person", "github", "followers")
	fmt.Println(string(rjson))
}
