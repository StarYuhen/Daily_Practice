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
	// res := rdb.Unlink(ctx, "748793fd-fb3d-4d7e-9f9c-d81427d6a08f")
	// logrus.Println(res)

	// 向哈希数据库插入数据
	// res, _ := rdb.HSet(ctx, "wdnmd", "keys---", "namess").Result()
	// logrus.Info(res)
	// 向hash数据库读取数据
	rest, _ := rdb.HMGet(ctx, "sdfdsfds", "jwt").Result()
	logrus.Info(rest)
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoicm9vdCIsImV4cCI6MTY0NjkxOTEwNywiaXNzIjoiU3Rhcll1aGVuIn0.ElMmowz88DcWvqpnHl5QkRRRG0T0rlhJXogUJtP3DBw
	//	err := rdb.Set(ctx, "wdnmd", "string wdnmd", time.Second*60).Err()
	// logrus.Info(err)

	// logrus.Info(rdb.Get(ctx, "wdnmds").Val())
}

func main() {
	ExampleClient()
}
