package main

import (
	"fmt"
)

// 测试当条件同时满足时的情况
func main() {
	// 创建通道通信
	ch := make(chan int)
	// 发送通信数据
	go func() {
		for {
			ch <- 10
		}
	}()

	// select监听
	for {
		select {
		case i := <-ch:
			fmt.Println(i, "---case1")
		case <-ch:
			fmt.Println("---case2")
		default:
			fmt.Println("yyyyy---case3")
		}
	}

	// 经过实验，我们可以发现当多个条件同时满足时，select会随机执行
}
