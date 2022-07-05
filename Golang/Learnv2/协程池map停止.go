package main

import (
	"github.com/gogf/gf/os/grpool"
	"github.com/sirupsen/logrus"
	"time"
)

var Send = make(map[string]*grpool.Pool, 0)

func main() {
	pool := grpool.New(10)
	Send["name"] = pool
	for i := 0; i <= 8; i++ {
		v := i
		pool.Add(func() {
			seconds := v * 2
			time.Sleep(time.Duration(seconds) * time.Second)
			if v == 5 {
				logrus.Info("协程池内容", Send["name"])

			}
			logrus.Info("协程池内容", v)
		})
	}

	Send["name"].Close()
	logrus.Info("停止任务")
	time.Sleep(1000 * time.Second)
}
