package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	dir := "./gzFiles2/gids"
	exist, err := PathExists(dir)
	if err != nil {
		logrus.Error("查询文件夹是否存在出错-->", err)
		return
	}

	if !exist {
		// 创建文件夹
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			logrus.Error("创建文件夹失败--->", err)
		}
		logrus.Println("创建文件夹成功--->", dir)
	}
}

// PathExists 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
