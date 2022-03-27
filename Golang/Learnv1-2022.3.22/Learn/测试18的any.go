package main

import (
	"log"
)

func main() {
	// 测试返回数据
	log.Println(recover() != any(nil))
}
