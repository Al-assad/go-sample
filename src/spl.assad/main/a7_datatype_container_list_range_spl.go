package main

import (
	"container/list"
	"container/ring"
	"fmt"
)

// @desc container 标准包提供的 list（双链表）、ring（环形链表） 集合操作

func main() {
	//listTest()
	ringTest()
}

// container/list 双链表使用
func listTest() {
	// 定义 list
	alist := list.New()

	// 向 list 加入元素
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range arr {
		alist.PushBack(v)
	}
	// 获取 list 长度
	fmt.Println(alist.Len())
	// 获取第一个元素
	firstE := alist.Front()
	// 获取最后一个元素
	lastE := alist.Back()
	fmt.Println(firstE.Value, lastE.Value)

	// 从头遍历 list
	for p := alist.Front(); p != nil; p = p.Next() {
		fmt.Print(p.Value, " ")
	}
	// 1 2 3 4 5 6 7 8 9
	println()
	// 从尾部遍历 list
	for p := alist.Back(); p != nil; p = p.Prev() {
		fmt.Print(p.Value, " ")
	}
	// 9 8 7 6 5 4 3 2 1
	println()

	// 移动元素
	// 将第 3 个元素移动到末尾
	p1 := alist.Front().Next().Next()
	alist.MoveToBack(p1)
	printList(alist) // 1 2 4 5 6 7 8 9 3
	// 将第 2 个元素移动到开头
	p1 = alist.Front().Next()
	alist.MoveToFront(p1)
	printList(alist) // 2 1 4 5 6 7 8 9 3
	// 将第 2 个元素移动到第 4 个元素后
	alist.MoveAfter(firstE.Next(), firstE.Next().Next().Next())
	printList(alist) // 2 1 5 6 4 7 8 9 3

	// 中间插入元素
	// 向第 2 个元素后加入 233
	alist.InsertAfter(233, alist.Front().Next())
	printList(alist) // 2 1 233 5 6 4 7 8 9 3

	// 删除元素
	// 删除第 3 个元素
	alist.Remove(alist.Front().Next().Next())
	printList(alist) //2 1 5 6 4 7 8 9 3

	// 清空所有元素
	alist.Init()
	printList(alist)         //
	fmt.Println(alist.Len()) // 0
}

func printList(alist *list.List) {
	for p := alist.Front(); p != nil; p = p.Next() {
		fmt.Print(p.Value, " ")
	}
	println()
}

// container/ring 环形链表使用
func ringTest() {
	// 定义 ring
	aring := ring.New(5)

	// 向 ring 添加元素
	arr := []int{1, 2, 3, 4, 5}
	for _, e := range arr {
		aring.Value = e
		aring = aring.Next()
	}

	// 遍历 ring
	p := aring
	for i := 0; i < aring.Len(); i++ {
		fmt.Print(p.Value, " ")
		p = p.Next()
	}
	println() // 1 2 3 4 5
	// 超出边界，重新回到起始指针
	for i := 0; i < 10; i++ {
		fmt.Print(p.Value, " ")
		p = p.Next()
	}
	println() // 1 2 3 4 5 1 2 3 4 5

	// 对 ring 中所有元素执行函数
	aring.Do(func(i interface{}) {
		fmt.Print(i, " ")
	})
	println() // 1 2 3 4 5

	// 快速移动到从 ring.Next() 起始的第 n % len 个元素，
	e := aring.Move(2)   // 获取第 3 个元素
	fmt.Println(e.Value) // 3

	// 在第 3 个元素后加入一个新链表 bring
	bring := ring.New(2)
	bring.Value = 71
	bring.Next().Value = 72
	aring.Move(2).Link(bring)
	printRing(aring) // 1 2 3 71 72 4 5

	// 删除从第 3 个元素开始的后 2 % len 个元素
	aring.Move(2).Unlink(2)
	printRing(aring) // 1 2 3 4 5
}

func printRing(aring *ring.Ring) {
	p := aring
	for i := 0; i < aring.Len(); i++ {
		fmt.Print(p.Value, " ")
		p = p.Next()
	}
	println()
}
