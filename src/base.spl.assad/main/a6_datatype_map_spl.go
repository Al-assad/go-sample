package main

import (
	"fmt"
	"sort"
)

// @desc Map 使用示例

func main() {
	//mapDefine()
	//mapSlice()
	sortedMap()
}

// Map 定义、赋值、使用
func mapDefine() {

	// 声明 map
	var dict1 map[string]int // or -  dict1 := make(map[string]int)
	// map 初始化
	dict1 = map[string]int{"one": 1, "two": 2}
	// 获取 map 值
	fmt.Println(dict1)
	fmt.Println(dict1["one"])
	// 修改 map 值
	dict1["one"] = 11
	fmt.Println(dict1["one"])
	// 新增 map 值
	dict1["three"] = 3
	fmt.Println(dict1["three"])

	// 声明后直接赋值
	dict2 := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(dict2)

	// 使用 make 声明柄创建空 map 指针
	dict3 := make(map[string]int)
	//dict3 := make(map[string]int, 100)  // 同时指定初始容量 cap
	dict3["one"] = 1
	dict3["two"] = 2
	dict3["three"] = 3

	// 判断元素是否存在
	_, ok := dict3["four"]
	fmt.Println(ok) // false
	// 在 if condition 中使用 map 元素存在判断
	if _, exist := dict3["three"]; exist {
		fmt.Println("dict3 exists key[\"three\"]")
	}

	// 删除元素
	delete(dict3, "two")

	// for-range 遍历 map
	for key, val := range dict3 {
		fmt.Printf("key=%s,val=%d \n", key, val)
	}
}

// 声明 map value 为 slice
func mapSlice() {
	mp := make(map[string][]string)
	mp["one"] = []string{"1", "one", "一"}
	mp["two"] = []string{"2", "two", "二"}
	for k, v := range mp {
		fmt.Print("key=", k, ",value=")
		for _, e := range v {
			fmt.Print(e, " ")
		}
		fmt.Println()
	}
	//key=one,value=1 one 一
	//key=two,value=2 two 二
}

// map 排序
// map 默认是无序的，不管是按照 key 还是按照 value 默认都不排序，
// 如果你想为 map 排序，需要将 key（或者 value）拷贝到一个切片，再对切片排序，
// 然后可以使用切片的 for-range 方法打印出所有的 key 和 value
func sortedMap() {
	// 按 key 排序
	barMap := map[string]int{
		"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
	for k, v := range barMap {
		fmt.Printf("%v=%v /", k, v)
	}
	// charlie=23 /golf=34 /alpha=34 /bravo=56 /foxtrot=12 /hotel=16 /indio=87 /juliet=65 /kili=43 /lima=98 /delta=87 /echo=56 /
	println()
	// 获取 key，装载到 slice
	keys := make([]string, len(barMap))
	i := 0
	for k, _ := range barMap {
		keys[i] = k
		i++
	}
	// key slice 排序
	sort.Strings(keys)
	// 打印
	for _, k := range keys {
		fmt.Printf("%v=%v /", k, barMap[k])
	}
	// alpha=34 /bravo=56 /charlie=23 /delta=87 /echo=56 /foxtrot=12 /golf=34 /hotel=16 /indio=87 /juliet=65 /kili=43 /lima=98 /
	println()
}
