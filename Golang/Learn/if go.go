package main

import "github.com/sirupsen/logrus"

func main() {
	logrus.Info(!(2 > 3 || 2 > 1))
}
