package main

import (
	"fmt"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/os/gtimer"
	"time"
)

func main() {
	gtimer.SetInterval(time.Second, func() {
		fmt.Println("SetInterval:", gtime.Now())
	})
	select {}
}
