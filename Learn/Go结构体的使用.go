package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type name struct {
	uid string
	us  string
}

func (name *name) cs() {
	name.uid = "13232"
	name.us = "wdsfds"
}

func (name *name) li() {
	fmt.Println(name.uid)
	fmt.Println(name.us)
}

func main() {
	logrus.Info(string([]byte{10}))
}
