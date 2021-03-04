package main

import (
	"fmt"
	"sync"
)

// @desc 使用 lock 互斥锁实现 Atomic 原子变量

type AtomicInt struct {
	lock sync.Mutex
	val  int
}

func (a *AtomicInt) Incre() {
	a.lock.Lock()
	a.val++
	a.lock.Unlock()
}

func (a *AtomicInt) Decre() {
	a.lock.Lock()
	a.val--
	a.lock.Unlock()
}

func (a *AtomicInt) GetVal() int {
	return a.val
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	mu := sync.Mutex{}
	atomVal := AtomicInt{mu, 0}

	go func(done func()) {
		defer done()
		atomVal.Incre()
	}(wg.Done)

	go func(done func()) {
		defer done()
		atomVal.Decre()
	}(wg.Done)

	atomVal.Incre()
	wg.Wait()
	fmt.Println(atomVal.GetVal()) //1
}
