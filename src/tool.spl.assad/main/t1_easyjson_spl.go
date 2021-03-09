package main

import (
	"fmt"
	"github.com/Al-assad/go-sample/src/tool.spl.assad/model"
	"time"
)

// @desc easyjson 高性能 json 库使用（固定结构 struct）
// easyjson 适用于固定结构 json 解析

/*
	go get -u github.com/mailru/easyjson
	go install github.com/mailru/easyjson/easyjson

	// 生成 easyjson 序列化/反序列化代码
	easyjson -all model/spl_json_model.go
*/

func main() {
	serializeJson()
	deserializeJson()
}

// 序列化
func serializeJson() {
	createtime, _ := time.Parse("2006-01-02", "2006-07-01")
	article := model.Article{
		Title:      "Bigtable: A Distributed Storage System for Structured Data",
		Auth:       model.Author{[]string{"Fay Chang", "Jeffrey Dean", "Sanjay Ghemawat", "Wilson C. Hsieh"}},
		CreateTime: createtime,
		Scope:      102408,
	}
	jsonString, err := article.MarshalJSON()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(jsonString))
	// {"title":"Bigtable: A Distributed Storage System for Structured Data","auth":{"name":["Fay Chang","Jeffrey Dean","Sanjay Ghemawat","Wilson C. Hsieh"]},"create_time":"2006-07-01T00:00:00Z","scope":102408}

}

// 反序列化
func deserializeJson() {
	jsonString := `{"title":"Bigtable: A Distributed Storage System for Structured Data","auth":{"name":["Fay Chang","Jeffrey Dean","Sanjay Ghemawat","Wilson C. Hsieh"]},"create_time":"2006-07-01T00:00:00Z","scope":102408}`
	var article model.Article
	err := article.UnmarshalJSON([]byte(jsonString))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(article)
	// {Bigtable: A Distributed Storage System for Structured Data {[Fay Chang Jeffrey Dean Sanjay Ghemawat Wilson C. Hsieh]} 2006-07-01 00:00:00 +0000 UTC 102408}
}
