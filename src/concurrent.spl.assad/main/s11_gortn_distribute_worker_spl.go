package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// @desc channel 实现任务分发的实例
/*
send task:  {1}
consume: {1 1 Worker-2}
send task:  {47}
send task:  {7}
consume: {7 21 Worker-4}
send task:  {59}
send task:  {1}
consume: {1 1 Worker-0}
send task:  {38}
consume: {38 63245986 Worker-2}
send task:  {25}
consume: {25 121393 Worker-4}
send task:  {60}
send task:  {56}
send task:  {20}
consume: {20 10946 Worker-4}
send task:  {54}
send task:  {31}
send task:  {2}
send task:  {49}
consume: {47 4807526976 Worker-3}
send task:  {8}
consume: {31 2178309 Worker-3}
*/

// 任务信息
type Task struct {
	FiboIndex int
}

// 任务结果
type TaskResult struct {
	FiboIndex  int
	FliboVal   int
	WorkerName string
}

func main() {
	pending := make(chan *Task, 3)
	done := make(chan *TaskResult)
	go sendTask(pending)
	// 分发 10 个处理协程
	for i := 0; i < 5; i++ {
		go Worker(pending, done, "Worker-"+strconv.Itoa(i))
	}
	go consumeTaskResult(done)

	select {}
}

// 发送任务，向 in 通道发送任务
func sendTask(in chan *Task) {
	for {
		task := &Task{rand.Intn(80)}
		in <- task
		fmt.Println("send task: ", *task)
		time.Sleep(5e8)
	}
}

// 工作协程，从 in 获取，处理结果写入 out
func Worker(in chan *Task, out chan *TaskResult, workerName string) {
	for {
		task := <-in
		result := &TaskResult{task.FiboIndex, fibonacci(task.FiboIndex), workerName}
		out <- result
	}
}

// 消费任务结果
func consumeTaskResult(out chan *TaskResult) {
	for {
		result := <-out
		fmt.Printf("consume: %v\n", *result)
	}
}

// 斐波那契数列计算
func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return res
}
