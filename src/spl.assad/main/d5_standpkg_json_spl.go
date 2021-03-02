package main

import (
	"encoding/json"
	"fmt"
)

// go json 处理（encoding/json）
// https://blog.golang.org/json

// JSON 与 Go 类型对应如下：
//    go bool：对应 json  booleans
//    go float64： 对应 json numbers
//    go string： 对应 json strings
//    go nil： 对应 json null
//    go []interface{}：对应 json array

type Person struct {
	Name     string
	Gender   int
	IsNewMan bool
}

func main() {
	// json 序列化
	person := &Person{"assad", 0, true}
	js, _ := json.Marshal(person)
	fmt.Println(string(js))
	// {"Name":"assad","Gender":0,"IsNewMan":true}

	// json 反序列化
	var person2 Person
	json.Unmarshal(js, &person2)
	fmt.Printf("Person %v\n", person2)
	// Person {assad 0 true}

	// 解码任意数据
	customJson := `{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`
	var st interface{}
	json.Unmarshal([]byte(customJson), &st) // json 字符串解码
	stMap := st.(map[string]interface{})    // 类型断言为 map[string] interface{}

	// 访问 json 节点，使用类型断言
	name := stMap["Name"].(string)
	age := int(stMap["Age"].(float64))
	parentsOri := stMap["Parents"].([]interface{})

	parents := make([]string, len(parentsOri))
	for i := 0; i < len(parentsOri); i++ {
		parents[i] = parentsOri[i].(string)
	}

	fmt.Println(name, age, parents)
	// Wednesday 6 [Gomez Morticia]

}
