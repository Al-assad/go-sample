package main

// @desc 实现惰性生成器

func main() {

}

// 斐波那契数列计算
func fibonacci2(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return res
}
