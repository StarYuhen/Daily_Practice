package main

import (
	"github.com/sirupsen/logrus"
	"github.com/staryuhen/FilesAPI"
	"os"
)

func main() {
	// 文件路径
	path, err := os.Getwd()
	if err != nil {
		logrus.Error("获取当前路径失败")
	}
	byteBuffer := FilesAPI.FileReadTxtForMat(path + "/Learn/1.txt")
	logrus.Info(string(byteBuffer))
}
