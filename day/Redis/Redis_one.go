package main

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var ctx = context.Background()

var rdb = sc()

func sc() *redis.Client {
	return redis.NewClient(&redis.Options{
		Network:      "tcp",
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,   // redis数据库index
		PoolSize:     100, // redis链接池，默认是4倍cpu数，这里固定 用于协程链接
		MinIdleConns: 50,  // 初始规定的redis，维护，让其不少于这个数
	})
}

func ExampleClient() {
	/*
		rdb := redis.NewClient(&redis.Options{
			Addr:       "localhost:6379",
			Password:   "", // no password set
			DB:         0,  // use default DB
			MaxRetries: 2000,
		})
	*/

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
	// 获取是否有这个集合
	res, err := rdb.SIsMember(ctx, "staryuhen", "14166760605xgdds").Result()
	if err != redis.Nil {
		logrus.Error("redis插入集合报错--->", err)
	}
	logrus.Println(res)

}

func main() {
	ExampleClient()
}
