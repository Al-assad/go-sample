package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// @desc go Error 错误处理

// 自定义异常结构
type UnsuportOperatorErr struct {
	Oper string
	Num1 float64
	Num2 float64
}
type NumCovErr struct {
	Num string
	Msg string
}

// 自定义异常需要继承 Error 方法
func (e UnsuportOperatorErr) Error() string {
	return fmt.Sprintf("Oper=%s, Num1=%f, Num2=%f", e.Oper, e.Num1, e.Num2)
}

func (e NumCovErr) Error() string {
	return fmt.Sprintf("num=%s, Msg=%s", e.Num, e.Msg)
}

func main() {

	safeCompute("20 + 30")
	safeCompute("")
	safeCompute("a + b")
	safeCompute("32 % 12")

}

// 异常捕获示例
func safeCompute(algo string) {
	r, err := computeAlgo(algo)
	if err == nil {
		fmt.Println(r)
		return
	}
	// 直接打印异常
	fmt.Printf("%s\n", err)
	// 等同于
	// fmt.Println(err.Error())

	// 根据不同异常类型处理
	switch err := err.(type) {
	case NumCovErr:
		fmt.Println("Num Convert Error:", err.Error())
	case UnsuportOperatorErr:
		fmt.Println("Unsupported Operator: ", err.Error())
	}
}

// 计算式求值，演示异常抛出
// algo: 1 + 2，20 / 10，40 - 1，50 * 2，etc
func computeAlgo(algo string) (float64, error) {
	if algo == "" {
		return 0, errors.New("algo is empty") // 创建异常
	}

	m := strings.Split(algo, " ")
	if len(m) != 3 {
		return 0, errors.New("algo format error") // 创建异常
	}

	a, aErr := strconv.ParseFloat(m[0], 64)
	if aErr != nil {
		return 0, NumCovErr{m[0], aErr.Error()} // 返回自定义异常
	}
	o := m[1]
	b, bErr := strconv.ParseFloat(m[2], 64)
	if bErr != nil {
		return 0, NumCovErr{m[2], bErr.Error()} // 返回自定义异常
	}

	switch o {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, UnsuportOperatorErr{o, a, b} // 返回自定义异常
	}
}
