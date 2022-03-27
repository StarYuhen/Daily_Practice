package main

import (
	"github.com/panjf2000/ants"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	var i = 0
	p, _ := ants.NewPool(10000)
	for b := 0; b <= 100; b++ {
		p.Submit(func() {
			time.Sleep(200)
			i++
			logrus.Info(i)
			ants.Release()
		})
	}
	time.Sleep(20000000)
}
