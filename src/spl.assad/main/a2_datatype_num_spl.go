package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	数值类型测试
*/

func main() {
	//test8()
	//test9()
	//test10()
	//test11Random()
	test12Char()
}

// 数值转换
func test8() {
	var n1 int16 = 12

	// 向上转换，安全
	var n2 int32 = int32(n1)
	println(n2) // 12

	// 向下转换，不一定安全
	var n3 int8 = int8(n1)
	println(n3) // 12
	n1 = 233
	var n4 int8 = int8(n1)
	println(n4) // -23
}

// 格式化输出
func test9() {
	a := 233
	b := 10.0 / 3
	// 整数，十六进制
	println(a)                           //233
	fmt.Printf("%d, %x, %X \n", a, a, a) // 233, e9, E9
	// 浮点
	println(b)               //+3.333333e+000
	fmt.Printf("%f \n", b)   //3.333333
	fmt.Printf("%.1f \n", b) // 输出小数点后一位 3.3
}

// 复数
func test10() {
	var c1 complex128 = 10 + 24i // 10.24
	fmt.Printf("%v \n", c1)
	println(real(c1)) // 获取实数部分
	println(imag(c1)) // 获取虚数部分

	var c2 = 3 + 12i
	var c3 = c1 + c2
	fmt.Printf("%v \n", c3)
}

// 随机数
func test11Random() {
	a := rand.Int()
	println(a)
	// 限制边界
	b := rand.Intn(10)
	println(b)
	// 手动设置种子
	timeNano := int64(time.Now().Nanosecond())
	rand.Seed(timeNano)
	println(rand.Int31())
}

// 字符类型
func test12Char() {
	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point

	// 65 - 946 - 1053236
	// A - β - r
	// 41 - 3B2 - 101234
	// U+0041 - U+03B2 - U+101234
}
