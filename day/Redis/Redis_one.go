package main

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var ctx = context.Background()

func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:       "localhost:6379",
		Password:   "", // no password set
		DB:         0,  // use default DB
		MaxRetries: 2000,
	})
	/*
	   	err := rdb.Set(ctx, "key", "value", 0).Err()
	   	if err != nil {
	   		panic(err)
	   	}

	   	val, err := rdb.Get(ctx, "key").Result()
	   	if err != nil {
	   		panic(err)
	   	}
	   	fmt.Println("key", val)

	   	val2, err := rdb.Get(ctx, "key2").Result()
	   	if err == redis.Nil {
	   		fmt.Println("key2 does not exist")
	   	} else if err != nil {
	   		panic(err)
	   	} else {
	   		fmt.Println("key2", val2)
	   	}
	   	// Output: key value
	   	// key2 does not exist
	   	csd, _ :=rdb.SAdd(ctx,"set add init staryuhen ","asfsasf").Result()
	   //	ERR, _ :=rdb.SRem(ctx,"测试当亲啊集合").Result()
	   	fmt.Println("是否元素中的所有元素",csd)
	   	s:=fmt.Sprintf("name\n")
	   	logrus.Info(s)
	*/

	res, err := rdb.SAdd(ctx, "redis 储存sgsd手dsgsddgds机号", "sdgddxcvxfgd", "dgdfhdfh").Result()
	if err != redis.Nil {
		logrus.Error("redis插入集合报错--->", err)
	}
	logrus.Println(res)

}

func main() {
	ExampleClient()
}
