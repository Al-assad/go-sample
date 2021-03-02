package main

import (
	"fmt"
	"time"
)

// @desc 时间日期类型使用

func main() {
	// 获取当前时间
	var now time.Time = time.Now()
	println(now.String()) // 2021-02-26 18:47:22.032675 +0800 CST m=+0.000156852
	// 获取时间戳
	println(now.Nanosecond())
	// 获取时间细节
	println(now.Year(), " ", now.Month(), " ", now.Day(), " ", now.Hour(), " ", now.Minute(), " ", now.Second(), " ", now.Weekday())

	// time 转化字符串
	fmt.Printf("%02d-%02d-%02d \n", now.Year(), now.Month(), now.Day())
	println(now.Format(time.RFC3339))
	println(now.Format("2006-01-02 15:04:05"))

	// 字符串构建 time
	time2 := time.Date(2020, 1, 1, 12, 12, 0, 0, time.Local)
	println(time2.String())
	time3, _ := time.Parse("2006-01-02 15:04:05", "2021-05-01 12:30:00")
	println(time3.String())

	// 比较时间
	println(time2.Before(time3))

	// 时间偏移
	time2.AddDate(0, 1, 0) //增加一个月偏移
	println(time2.String())
	time2.Add(time.Second * 102534) //增加 102534 秒偏移
	println(time2.String())

}
