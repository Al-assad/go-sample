package main

import (
	"bytes"
	"fmt"
	"math/rand"
)

// @desc 数组，切片（slice）使用示例

func main() {
	//arrayTest()
	//arrayCopy()
	//arrayConst()
	// matrixTest()
	//transferArrToFunc()
	//sliceDefine()
	//reslice()
	//sliceCopyAppend()
	sliceUnknownCap()
	//bufferTest()
}

// 数组定义、赋值、遍历
func arrayTest() {
	// 声明
	var arr1 [5]string // or -  arr1 := new([5]string)
	// 赋值
	arr1[0] = "are"
	arr1[1] = "you"
	arr1[2] = "ok"
	// 获取长度
	fmt.Println(len(arr1))
	// 遍历 - 1
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("%d:%s \n", i, arr1[i])
	}
	// 遍历 - 2
	for i, e := range arr1 {
		fmt.Printf("%d:%s \n", i, e)
	}
}

// 数组拷贝，默认的值拷贝，和引用拷贝方式
func arrayCopy() {
	arr1 := [3]int{1, 2, 3}
	// go 数组默认是值拷贝
	arr2 := arr1 // arr2 是对 arr1 的值拷贝，对 arr2 的修改不会传递到 arr1
	arr2[0] = 233
	fmt.Println(arr1, arr2) // [1 2 3] [233 2 3]

	// 引用拷贝数组
	arr3 := &arr1 // 数组 arr3 是 arr1 的指针引用，对 arr3 的修改会传递到 arr1
	arr3[0] = 233
	fmt.Println(arr1, arr3) // [233 2 3] &[233 2 3]
}

// 数组常量，声明数组时进行赋值
func arrayConst() {
	var arr1 = [3]int{1, 2, 3}
	var arr2 = [10]int{1, 2, 3} // arr2 10 个元素，从第4个元素开始都是默认值0

	arr3 := [...]int{1, 2, 3} // 根据初始化元素自动确定数组长度
	arr4 := []int{1, 2, 3}

	arr5 := [5]string{3: "are", 4: "you"} // 指定元素及其索引值

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5, arr5[0] == "")
}

// 多维数组
func matrixTest() {
	var pixel [10][10]int

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			pixel[x][y] = 6
		}
	}
	fmt.Println(pixel)
}

// 向函数传递数组
func transferArrToFunc() {
	arr := [20]int{1, 2, 3, 4, 5}
	arrSize := count(&arr) // 传递数组指针
	fmt.Println(arrSize)

	slice := arr[:]
	arrSize2 := count2(slice) // 传递数组分片
	fmt.Println(arrSize2)
}

// 向函数传递数组：传递数组的指针
func count(arr *[20]int) int {
	index := 0
	for _, v := range *arr {
		if v == 0 {
			return index
		}
		index++
	}
	return index
}

// 向函数传递分片
func count2(arr []int) int {
	index := 0
	for _, v := range arr {
		if v == 0 {
			return index
		}
		index++
	}
	return index
}

// slice 切片定义
func sliceDefine() {
	// 创建指向 arr 切片
	arr := []string{"a", "b", "c", "d", "e"}
	slice1 := arr[:] // 等同于 slice1 := &arr
	slice2 := arr[1:3]
	slice3 := arr[4:]
	fmt.Printf("%s, cap=%d, len=%d \n", slice1, cap(slice1), len(slice1)) // [a b c d e], cap=5, len=5
	fmt.Printf("%s, cap=%d, len=%d \n", slice2, cap(slice2), len(slice2)) // [b c], cap=4, len=2
	fmt.Printf("%s, cap=%d, len=%d \n", slice3, cap(slice3), len(slice3)) // [e], cap=1, len=1

	// 直接创建分片
	slice4 := make([]int, 10, 100) // cap=10, len=100
	println(len(slice4), cap(slice4))
	for i := 0; i < len(slice4); i++ {
		slice4[i] = rand.Int()
	}
	fmt.Println("cap=%d, len=%d, %s \n", cap(slice4), len(slice4), slice4)

	// 直接创建切片，仅指定 cap
	s2 := make([]int, 100)        // 指定 cap
	fmt.Println(len(s2), cap(s2)) // 100, 100
}

// slice 切片重组（reslice）
func reslice() {
	s := make([]int, 10, 100) // 指定 len、cap
	for i := 0; i < len(s); i++ {
		s[i] = 7
	}
	fmt.Printf("len=%d, cap=%d, s[9]=%d \n", len(s), cap(s), s[9]) // len=10, cap=100, s[9]=7
	// slice 右边界向左缩容 5 位
	s = s[0 : len(s)-5]
	fmt.Printf("len=%d, cap=%d \n", len(s), cap(s)) // len=5, cap=100
	// slice 右边界向后扩容 15 位
	s = s[0 : len(s)+15]
	fmt.Printf("len=%d, cap=%d, s[9]=%d \n", len(s), cap(s), s[9]) //len=20, cap=100, s[9]=7

	// 可以看到，slice 由于本身只是一个指针，缩容并不影响原数组元素值
}

// slice 复制、追加
func sliceCopyAppend() {
	// slice 复制
	s1 := []int{1, 2, 3}
	s2 := make([]int, 10)
	copy(s2, s1)    // 把 s1 复制到 s2，可以使用这种方式返回部分数组值拷贝，释放整个原数组内存资源
	fmt.Println(s2) // [1 2 3 0 0 0 0 0 0 0]

	// slice 追加
	s3 := append(s1, 4, 5, 6)
	fmt.Println(s3) // [1 2 3 4 5 6]

}

// 在不确定 slice 长度的情况下，可以用 append 的方式扩容 slice 长度
func sliceUnknownCap() {
	// slice 追加多个元素
	s1 := []int{1, 2, 3}
	s1 = append(s1, 4, 5, 6)
	fmt.Println(s1)
	// [1 2 3 4 5 6]

	// 将 slice 所有元素一次性追加另一个 slice
	s2 := []int{1, 2, 3}
	s3 := []int{4, 5, 6}
	s2 = append(s2, s3...)
	fmt.Println(s2)
	// [1 2 3 4 5 6]
}

// Buffer 使用，处理 byte[]
func bufferTest() {
	content := "are you ok? 你还好吗？"
	contentRune := []rune(content)
	var buffer bytes.Buffer
	for i := 0; i < len(contentRune); i++ {
		buffer.WriteString(string(contentRune[i]))
	}
	fmt.Println(buffer.String())
}
