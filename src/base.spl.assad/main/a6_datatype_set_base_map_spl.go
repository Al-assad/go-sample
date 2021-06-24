package main

import "fmt"

// @desc 基于 map 的 set 结构实现

// go 没有 set 结构，常常使用 map[interface{}] struct{} 来代替实现

type StrSet map[string]struct{}

func (set StrSet) Has(e string) bool {
	_, ok := set[e]
	return ok
}

func (set StrSet) Add(e string) {
	set[e] = struct{}{}
}

func (set StrSet) Delete(e string) {
	delete(set, e)
}

func (set StrSet) Iterate() <-chan string {
	ch := make(chan string)
	go func() {
		for key := range set {
			ch <- key
		}
		close(ch)
	}()
	return ch
}

func main() {
	set := make(StrSet)
	set.Add("guangzhou")
	set.Add("shanghai")
	set.Add("beijing")

	fmt.Println(set.Has("guangzhou"))
	set.Delete("guangzhou")
	fmt.Println(set.Has("guangzhou"))

	// 遍历
	for e, _ := range set {
		fmt.Println(e)
	}

	// 遍历2
	for e := range set.Iterate() {
		fmt.Println(e)
	}

}
