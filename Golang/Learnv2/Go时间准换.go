package main

import (
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	logrus.Info(24 * 7 * 60 * 60)
	logrus.Info(time.Now().Unix() - 24*7*60*60)
	logrus.Info(time.Now().Unix())
}
