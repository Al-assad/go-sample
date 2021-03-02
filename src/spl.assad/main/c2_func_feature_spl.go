package main

import "time"

// 函数特性示例

func main() {
	println(add(3, 4))
	am, user := say("assad")
	println(am, user)

	// 获取多返回值
	name, greet := say2("assad")
	println(greet, name)

	// 空白符丢弃
	_, greet2 := say2("assad")
	println(greet2)

	println(sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9))

	foo()
}

// 普通函数
func add(a, b int) int {
	return a + b
}

// 多返回值
func say(user string) (int, string) {
	am := 0
	if time.Now().Hour() >= 12 {
		am = 1
	}
	return am, user
}

// 多返回值，返回值命名
func say2(user string) (name string, greetContent string) {
	return user, "Hello!"
}

// 传递可变长参数
func sumAll(num ...int) int {
	sum := 0
	for _, n := range num {
		sum += n
	}
	return sum
}

// defer 最终执行，类似 Java finally
func foo() {
	println("into foo()")
	defer println("after out foo()") // 在 foo return 后执行
	println("out foo()")
	//into foo()
	//out foo()
	//after out foo()
}
