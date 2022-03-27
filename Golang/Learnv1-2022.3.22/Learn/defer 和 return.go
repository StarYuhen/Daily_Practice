package main

import "fmt"

//测试defer和return的碰撞
func main() {
	sd()
}

func sd() int {
	fmt.Println("我是开头")
	defer func() { fmt.Println("我是defer") }()
	return 100
}
