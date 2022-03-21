package main

import (
	"fmt"
	"sync"
)

var a sync.Mutex
var b sync.RWMutex

// Go倘若没有用到协程就不必加锁了

func main() {
	for i := 1; i <= 100; i++ {
		/*
			go func() {
				a.Lock()
				fmt.Println("测试锁",i)
				//defer a.Unlock()
			}()

		*/

		go func() {
			b.Lock()
			fmt.Println("测试机", i)
			//defer b.Unlock()
		}()

	}
}
