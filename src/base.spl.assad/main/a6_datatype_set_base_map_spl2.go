package main

import "fmt"

// @desc 基于 map 的 set 结构实现

type Set struct {
	inner map[interface{}]struct{}
}

func NewSet() *Set {
	set := new(Set)
	set.inner = make(map[interface{}]struct{})
	return set
}

func (set *Set) Has(e interface{}) bool {
	_, ok := set.inner[e]
	return ok
}

func (set *Set) Add(e interface{}) {
	set.inner[e] = struct{}{}
}

func (set *Set) Delete(e interface{}) {
	delete(set.inner, e)
}

func (set *Set) Size() int {
	return len(set.inner)
}

func (set *Set) Iterate() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		for key := range set.inner {
			ch <- key
		}
		close(ch)
	}()
	return ch
}

func main() {
	set := NewSet()
	set.Add(12)
	set.Add(13)
	set.Add(14)
	fmt.Println(set.Has(12))
	for e := range set.Iterate() {
		fmt.Println(e)
	}

}
