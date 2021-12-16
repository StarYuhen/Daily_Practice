package main

import (
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	timenow := time.Now().AddDate(0, 0, -1).Format("20060102")
	// timews := strings.Split(timenow, " ")
	logrus.Info(timenow)
}
