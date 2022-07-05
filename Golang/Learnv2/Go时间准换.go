package main

import (
	"fmt"
	"time"
)

func main() {
	// str := "2022/05/19 22:36:12"
	// loc, _ := time.LoadLocation("Local")
	// res, _ := time.ParseInLocation("2006-01-02 15:04:05", str, loc)
	//
	// fmt.Println(res)

	// local, _ := time.LoadLocation("Local")
	//
	// t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-05-20T12:54:08.5659852+08:00", local)
	ts, _ := time.Parse(time.RFC3339, "2022-05-20T12:54:08.5659852+08:00")
	fmt.Println(ts.Unix())
	fmt.Println(time.Now())
}
