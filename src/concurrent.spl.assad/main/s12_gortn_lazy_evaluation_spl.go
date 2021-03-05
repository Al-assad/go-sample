package main

import "fmt"

// @desc 实现惰性生成器

func main() {
	sequence := NewEvenSequence()
	fmt.Println(sequence.Next()) //0
	fmt.Println(sequence.Next()) //2
	fmt.Println(sequence.Next()) //4

}

// 无限获取的偶数数列
type EvenSequence struct {
	resume chan int // 当前偶数值
}

// 构建函数
func NewEvenSequence() *EvenSequence {
	ch := genEven()
	sequence := &EvenSequence{ch}
	return sequence
}

// 获取下一个偶数
func (e *EvenSequence) Next() int {
	return <-e.resume
}

// 构建偶数协程，通过 channel 实现类似 yield 机制
func genEven() chan int {
	yield := make(chan int)
	curVal := 0
	go func() {
		for {
			yield <- curVal
			curVal += 2
		}
	}()
	return yield
}
