package func_foo

/*
	函数测试
*/

// 包内可访问函数
func minus(a, b int) int {
	return a - b
}

// 包外可访问函数
func Minus(a int, b int) int {
	return a - b
}
