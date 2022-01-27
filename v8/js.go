package main

import (
	"github.com/dop251/goja"
	"github.com/sirupsen/logrus"
)

func main() {
	// 启用js虚拟机
	vm := goja.New()
	// 运行js代码
	v, err := vm.RunString("2 + 2")
	if err != nil {
		panic(any(err))
	}
	if num := v.Export().(int64); num != 4 {
		panic(any(num))
	}
	logrus.Info(v)
}
