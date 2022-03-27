package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

// 定时器 鸡肋

func main() {
	times, _ := time.Parse("2006-01-02 15:04:05", "2014-06-15 08:37:18")
	timeUnix := times.Unix()
	fmt.Printf("times is %+v \n, timeUnix is %+v", times, timeUnix)
	logrus.Info(-10 > -100)
}
