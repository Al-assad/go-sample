package main

// 包内可访问函数
func sum(a, b int) int {
	return a + b
}

// 包外可访问函数
func Sum(a int, b int) int {
	return a + b
}
