package main

import (
	"github.com/panjf2000/ants"
	"github.com/sirupsen/logrus"
)

var SendTasks map[string]*ants.Pool

func init() {
	SendTasks = make(map[string]*ants.Pool, 0)
}

func main() {
	logrus.Info(SendTasks["name"] == nil)
}
