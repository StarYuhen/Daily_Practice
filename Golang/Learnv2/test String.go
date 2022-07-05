package main

import (
	"github.com/sirupsen/logrus"
	"strings"
)

func main() {
	ls := strings.Split("name", "|")
	logrus.Info(len(ls))
}
