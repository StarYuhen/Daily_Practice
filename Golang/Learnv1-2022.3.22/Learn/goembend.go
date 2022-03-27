package main

import (
	_ "embed"
	"github.com/sirupsen/logrus"
)

// 使用go:embed 进行本地绑定文件内容到变量

//go:embed 1.txt
var str string

func main() {
	logrus.Info(str)
}
