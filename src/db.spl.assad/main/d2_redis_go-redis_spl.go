package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// @desc redis 操作（go-redis/redis）

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(
		&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		})

	// 新增数据
	err := rdb.Set(ctx, "key123", "helloworld", 10*time.Minute).Err()
	if err != nil {
		panic(err)
	}

	// 查询数据
	val, err := rdb.Get(ctx, "key123").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

	// 查询数据
	val2, err := rdb.Get(ctx, "key123").Result()
	if err == redis.Nil {
		fmt.Println("key123 is not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println(val2)
	}

}
