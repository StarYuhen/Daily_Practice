package main

import "github.com/sirupsen/logrus"

func main() {
	s := "\n"
	for index, val := range s {
		logrus.Info(index, val, string(val))
	}
}
