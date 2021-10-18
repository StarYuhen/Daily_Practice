package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
)

// redis 函数和文档 https://www.yuhenm.com/index.php/archives/215/

var ct = context.Background()

func main() {
	// 连接redis
	redisDB := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // 数据库地址
		Password: "",               // 数据库密码
		DB:       0,                // 数据库连接数
		/*   启用https
		TLSConfig: &tls.Config{
			ServerName: "www.yuhenm.com",
		},
		*/
	})

	// age是key,19是value,30000是这个值的过期时间，倘若是0就是永不过期 增加这么一个值
	set := redisDB.Set(ct, "age", "19", 30000)
	fmt.Println(set)
	RDB("age", redisDB)

	// 向list添加元素
	redisDB.RPush(ct, "list", "wdnmd")
	length, err := redisDB.LLen(ct, "list").Result()
	if err == redis.Nil {
		panic(err)
	}
	fmt.Println("list的长度", length)
	fmt.Println(redisDB.Incr(ct, "WDNMS"))
}

func RDB(key string, redisDB *redis.Client) {
	val, err := redisDB.Get(ct, key).Result()
	// redis.Nil指的是空字符串或者nil
	if err == redis.Nil {
		fmt.Println("key not")
	}
	// 访问其对应的value值
	fmt.Println(val)
}
