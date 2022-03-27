package main

import (
	"fmt"
	"github.com/gogf/gf/os/gtimer"
	"time"
)

func main() {
	gtimer.AddOnce(time.Second, func() {
		fmt.Println("sdfsd")
	})

	time.Sleep(time.Hour)
}
