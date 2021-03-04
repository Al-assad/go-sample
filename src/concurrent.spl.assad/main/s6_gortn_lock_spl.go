package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// @desc 使用 lock 互斥锁实现公有数据的互斥访问

var (
	wg      sync.WaitGroup
	num     = 0
	safeNum = &SafeNum{sync.Mutex{}, 0}
)

// 共有变量持有互斥锁
type SafeNum struct {
	MuLock sync.Mutex // 互斥锁
	Num    int
}

func main() {
	dataVisitedCompetitionSpl()
	safeDataVisitedSpl()
}

// 数据竞争的例子
func dataVisitedCompetitionSpl() {
	wg.Add(2)
	go updateNum(10, wg.Done)
	go updateNum(10, wg.Done)
	wg.Wait()
	fmt.Println(num) // 10
}

func updateNum(update int, done func()) {
	defer done()
	cur := num
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	num = cur + update
}

// 使用互斥锁解决数据安全访问
func safeDataVisitedSpl() {
	wg.Add(2)
	go safeUpdateNum(10, wg.Done)
	go safeUpdateNum(10, wg.Done)
	wg.Wait()
	fmt.Println(safeNum.Num) // 20
}

func safeUpdateNum(update int, done func()) {
	defer done()
	safeNum.MuLock.Lock() // 互斥加锁
	cur := safeNum.Num
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	cur += update
	safeNum.Num = cur
	safeNum.MuLock.Unlock() // 互斥锁解锁
}
