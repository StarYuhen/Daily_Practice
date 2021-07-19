package main

import (
	"fmt"
	"time"
)

//定时器 鸡肋

func main() {
	time.AfterFunc(200, func() {
		fmt.Println("定时器")
	})
	time.Sleep(200000000000)
}
