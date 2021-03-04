package func_foo

import (
	"math/rand"
	"testing"
)

// @desc go 单元测试、基准测试
// 测试 foo.go
// go test 指令工具提供了单元测试执行，可以对一个包进行单元测试，会同时编译包中的代码和测试代码
/*
可以使用以下函数来通知测试失败
	1）func (t *T) Fail() ： 标记测试函数为失败，然后继续执行（剩下的测试）
	2）func (t *T) FailNow() ：标记测试函数为失败并中止执行；文件中别的测试也被略过，继续执行下一个文件
	3）func (t *T) Log(args ...interface{}) ：args 被用默认的格式格式化并打印到错误日志中
	4）func (t *T) Fatal(args ...interface{}) ：结合 先执行 3），然后执行 2）的效果
*/

// 单元测试 Minus
func TestMinus(t *testing.T) {
	if r := Minus(2, 1); r != 1 {
		t.Log("2-1 must be 1")
		t.Fail() //标记为失败，并继续执行以下测试
	}
	if r := Minus(10, 6); r != 4 {
		t.Fatal("10-6 must be 4")
		// 等同于
		//t.Log("10-6 must be 4")
		//t.FailNow()
	}
	/*
		=== RUN   TestMinus
		--- PASS: TestMinus (0.00s)
		PASS
	*/
}

// 基准测试 Minus
// func 必须为 Benchmarkxx
func BenchmarkMinus(b *testing.B) {
	Minus(rand.Int(), rand.Int())
	/**
	goos: darwin
	goarch: amd64
	pkg: spl.assad/func_foo
	BenchmarkMinus
	BenchmarkMinus-12    	1000000000	         0.000000 ns/op
	PASS
	*/
}

// 批量测试数据，期望值
var minusTestData = []struct {
	in1 int
	in2 int
	out int
}{
	{10, 5, 5},
	{10000, 5000, 5000},
	{0, 5000, -5000},
	{7, 7, 0},
}

// 批量测试
func TestMinus2(t *testing.T) {
	for i, data := range minusTestData {
		r := Minus(data.in1, data.in2)
		if data.out != r {
			t.Errorf("%d. %q => %q , wanted: %q", i, data, r, data.out)
		}
	}
	/*
		=== RUN   TestMinus2
		--- PASS: TestMinus2 (0.00s)
		PASS
	*/
}
