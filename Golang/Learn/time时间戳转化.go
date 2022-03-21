package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	/*
		str:=time.Now()
		log.Println(str)
		year := time.Now().Year() // 年

		fmt.Println(year)

		month := time.Now().Month() // 月
		fmt.Println(month)

		day := time.Now().Day() // 日
		fmt.Println(day)

		hour := time.Now().Hour() // 小时
		fmt.Println(hour)

		minute := time.Now().Minute() // 分钟
		fmt.Println(minute)

		second := time.Now().Second() // 秒
		fmt.Println(second)

		nanosecond := time.Now().Nanosecond() // 纳秒
		fmt.Println(nanosecond)


	*/

	var ints int64 = 86400
	t := time.Now().Unix()
	s := time.Unix(t, 0)
	logrus.Info(1635987779 + ints*2)
	if t > 1635987779+ints*2 {
		fmt.Println("账号注册时间符合要求")
	}
	fmt.Println(s.Format("2006/01/02 15:04:05"))
	fmt.Println(t)

}
