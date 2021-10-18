package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	// string函数实践
	var cty = context.Background()
	// 链接数据库
	clim := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// string操作 ---------------------------------
	fmt.Println("string操作 ---------------------------------")
	// 为名字是user的string的value赋值   3000是过期时间，0则是永远不过期
	fmt.Println("为名字是user的string的value赋值", clim.Set(cty, "user", "admin", 3000))
	fmt.Println("为名字是age的string的value赋值", clim.Set(cty, "age", 0, 3000))

	/*
		数字的自增和自减 通常时点赞和评论数字
	*/
	fmt.Println("让age自增", clim.Incr(cty, "age"))
	fmt.Println("让age自减", clim.Decr(cty, "age"))

	// 返回数值为user的string的value值
	fmt.Println("返回数值为user的string的value值", clim.Get(cty, "user"))

	// 返回库中多个key的value
	fmt.Println("返回库中多个key的value", clim.MGet(cty, "user", "age"))

	// 添加key和value
	fmt.Println("添加key和value", clim.SetNX(cty, "string_user", "18h", 3000))

	// 向string_user的value后面添加wdnmd
	fmt.Println("向key的value后面添加内容", clim.Append(cty, "string_user", "wdnmd"))
	// 最后结果为 18hwdnmd
	fmt.Println("添加的最后结果为", clim.Get(cty, "string_user"))

	// list/切片 操作    -----------------------------------
	fmt.Println("list/切片 操作    -----------------------------------")

	// 在key的list尾部添加元素
	fmt.Println("在key的list尾部添加元素", clim.RPush(cty, "list_user", "test"))

	// 在key头部添加元素
	fmt.Println("在key头部添加元素", clim.LPush(cty, "list_user", "append head"))

	// 返回key的长度
	fmt.Println("返回key的长度", clim.LLen(cty, "list_user"))

	// 返回指定list索引内容
	fmt.Println("返回指定list索引内容", clim.LIndex(cty, "list_user", 1))

	// 返回指定list索引内容范围 1-6
	fmt.Println("返回指定list索引内容范围 1-6", clim.LTrim(cty, "list_user", 1, 6))

	// 给list指定索引地址赋值
	fmt.Println("给list指定索引地址赋值", clim.LSet(cty, "list_user", 1, "wdnmd 我是赋值的"))

	// 给list指定索引地址删除内容
	fmt.Println("删除list的所有包含test的内容", clim.LRem(cty, "List_user", 2, "test"))

	// 删除list_user的最后一个元素
	fmt.Println("删除list_user的最后一个元素", clim.RPop(cty, "list_user"))

	// 删除list_user的第一个元素
	fmt.Println("删除list_user的第一个元素", clim.LPop(cty, "list_user"))

	// 集合set操作--------------------------
	fmt.Println("集合set操作--------------------------")

	// 向名称为set_use(他是key)r添加元素
	fmt.Println("向名称为set_user添加元素", clim.SAdd(cty, "set_user", "wdnmd", "sdfsdg"))

	// 删除名称为set_user的元素
	fmt.Println("删除名称为set_user的元素", clim.SRem(cty, "set_user", "wdnmd"))

	// 随机删除set_user中的一个元素
	fmt.Println("随机删除set_user中的一个元素", clim.SPop(cty, "set_user"))

	// 移动元素到另一个集合  没有元素会返回false
	fmt.Println("移动元素到另一个集合", clim.SMove(cty, "set_user", "set_user_in", "wdnmd"))
}
