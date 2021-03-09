package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// @desc 控制结构语句（if-else, for-i, for-range, for-condition, goto）

func main() {
	//ifElseTest()
	//switchTest()
	//forTest()
	gotoTest()
}

// if else 结构
func ifElseTest() {
	var flag = rand.Intn(2) > 0
	if flag {
		println("true")
	} else {
		println("false")
	}

	// 条件中赋值
	if val := rand.Int(); val%2 == 1 {
		println("Odd")
	} else {
		println("Even")
	}

	// 错误校验
	num, err := strconv.Atoi("2333.2")
	if err == nil {
		println(num)
	} else {
		println(err.Error())
	}

}

// switch 结构
func switchTest() {
	const (
		Monday = iota + 1
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)
	today := time.Now().Weekday()

	switch today {
	case Monday:
		println("sad!")
	case Tuesday:
	case Wednesday:
	case Thursday:
		println("just so so.")
	case Friday:
		println("on the eve of happy")
	default:
		println("partying all night")
	}
}

// go switch case 可以接收表达式，因此完全可以使用 switch 代替 if-elseif-else 的结构
func switchExpresion() {
	val := 233
	// 使用 switch 代替 if-elseif-else
	switch {
	case val >= 0 && val < 100:
		fmt.Println("in [0,100)")
	case val >= 100 && val < 200:
		fmt.Println("in [100,200)")
	case val >= 200 && val < 300:
		fmt.Println("in [200,300)")
	default:
		fmt.Println("out [-~,100) and [300, +~)")
	}

	// switch case 不支持条件下推，但是可以使用 "," 分隔多 case 条件
	str := "hello, assad"
	switch {
	case strings.Contains(str, "hello"), strings.Contains(str, "hi"), strings.Contains(str, "halo"):
		fmt.Println("Hi")
	default:
		fmt.Println("Unable to Understand!")
	}
}

// for 结构
func forTest() {
	// for-i
	for i := 0; i < 10; i++ {
		println("count: ", i)
	}

	// for-i
	poem := "Whenever you need me， I'll be here"
	words := strings.Split(poem, " ")
	for i := 0; i < len(words); i++ {
		println(words[i])
	}

	// for-range
	for i := range words {
		println(words[i])
	}
	for i, word := range words {
		println(i, word)
	}
	for _, word := range words {
		println(word)
	}

	// while
	for true {
		val := rand.Int()
		if val%2 == 0 {
			println(val)
			break
		}
	}
}

// goto 结构
func gotoTest() {
	i := 0
HERE:
	println(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}
