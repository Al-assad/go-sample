package main

import (
	"fmt"
	"time"
)

// @desc go 协程通讯 - channel 使用示例
/**

// 声明通道，int 类型数据
ch := make(chan int)
// 把数据 v 发送到 ch
ch <- v
// 从信道获取数据 v
v := <- ch

*/

func main() {
	testChannel()
	//testChannel2()
	//testChannel3()
}

// 通道使用示例 1 - 默认单个通道容量
func testChannel() {
	ch := make(chan string) // 创建通道，默认通道容量1
	go sendData(ch)
	go getData(ch)
	time.Sleep(10 * time.Second)
	/** 可以看到，由于通道容量 1，在消费协程消费通道数据之前，发送协程通道发送动作被阻塞
	send: Washington
	send: Tripoli
	receive:  Washington
	send: London
	receive:  Tripoli
	receive:  London
	send: Beijing
	receive:  Beijing
	send: Tokyo
	receive:  Tokyo
	*/
}

// 向通道发送数据
func sendData(ch chan string) {
	data := []string{"Washington", "Tripoli", "London", "Beijing", "Tokyo"}
	for _, da := range data {
		fmt.Println("send:", da)
		ch <- da // 向通道发送数据
		//time.Sleep(500 * time.Millisecond)
	}
}

// 从通道接收数据
func getData(ch chan string) {
	var input string
	for {
		input = <-ch // 从通道获取数据
		fmt.Println("receive: ", input)
		time.Sleep(1 * time.Second)
	}
}

// 通道使用示例 2 - 设置通道缓冲容量
func testChannel2() {
	ch := make(chan string, 10)
	go sendData(ch)
	go getData(ch)
	time.Sleep(10 * time.Second)
	/**
	send: Washington
	send: Tripoli
	send: London
	send: Beijing
	send: Tokyo
	receive:  Washington
	receive:  Tripoli
	receive:  London
	receive:  Beijing
	receive:  Tokyo
	*/
}

// 通道使用示例 3 - 无通道消费
func testChannel3() {
	ch := make(chan string)
	go sendData(ch)
	time.Sleep(10 * time.Second)
	/** 由于无通道数据消费，当通道容量写满后，其他协程写入通道动作被阻塞
	send: Washington
	*/
}
