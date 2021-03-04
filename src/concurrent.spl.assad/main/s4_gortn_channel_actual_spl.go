package main

import (
	"fmt"
	"math/rand"
	"time"
)

// @desc 一个协程 channel 实际使用示例（计算分发）

func main() {

	//r := splitSum(genRandInts(100000))
	//fmt.Println(r)

	testData := genRandInts(1000000000)
	r := blockSum(testData, 1000) // use： 3.122856472s
	fmt.Println(r)
	r2 := blockSum(testData, 100) // use：use： use： 7.883396615s
	fmt.Println(r2)
	r3 := blockSum(testData, 10) // use： use： 12.759250593s
	fmt.Println(r3)
	r4 := blockSum(testData, 1) // use： 16.136398053s
	fmt.Println(r4)
}

// Channel 使用示例，演示基本 Channel 的声明，数据的获取
// 模拟对半分块统计
func splitSum(nums []int) int {
	ch := make(chan int)                    // 声明通道
	go longCostSum(nums[0:len(nums)/2], ch) // 传递通道到协程函数
	go longCostSum(nums[len(nums)/2:], ch)

	r1, r2 := <-ch, <-ch // 从通道获取结果
	close(ch)            // 关闭通道
	return r1 + r2
}

// Channel 使用示例，演示定义 Channel 缓冲区
// 模拟分块统计
func blockSum(nums []int, chBufSize int) int {
	// 统计运行时间
	launch := time.Now()
	defer func() {
		delay := time.Now().Sub(launch)
		fmt.Println("use：", delay)
	}()

	ch := make(chan int, chBufSize) // 通道缓冲区容量

	const blockSize = 1000 // 分块容量
	// 分发计算任务
	taskCount := 0
	for i := 0; i < len(nums); i += blockSize {
		taskCount++
		if end := i + blockSize; end > len(nums) {
			go longCostSum(nums[i:], ch)
		} else {
			go longCostSum(nums[i:i+blockSize], ch)
		}
	}
	// 从通道中获取计算结果
	sum := 0
	for i := 0; i < taskCount; i++ {
		sum += <-ch
	}
	close(ch)
	return sum
}

// 模拟耗时的统计计算，演示 Channel 数据的发送
func longCostSum(nums []int, ch chan int) {
	time.Sleep(1 * time.Second)
	sum := 0
	for _, n := range nums {
		sum += n
	}
	ch <- sum // 将 sum 发送到 ch 通道
	time.Sleep(1 * time.Second)
}

func genRandInts(size int) []int {
	s := make([]int, size)
	for i := 0; i < size; i++ {
		s[i] = rand.Intn(100)
	}
	return s
}
