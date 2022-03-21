package main

import (
	"fmt"
	"time"
)

var access = 1

//协程能更改变量内容，无须传递

func main() {

	go func() {
		for p := 1; p <= 100; p++ {
			access = 2
		}
	}()

	go func() {
		for p := 1; p <= 100; p++ {
			access = 3
		}
	}()

	for p := 1; p <= 300; p++ {
		fmt.Println(access)
	}

	b := 10

	go func() {
		b = 100
	}()

	time.Sleep(2000)

	//可以得知协程可以更改变量内容
	fmt.Println(b)

}
