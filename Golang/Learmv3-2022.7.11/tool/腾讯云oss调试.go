package main

import (
	"BackEnd/config"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"strings"
)

func main() {
	// tx:= config.TxOss()
	// // 获取账号所有储存桶
	// list, _, _ := tx.Service.Get(context.Background())
	// for _, v := range list.Buckets {
	// 	logrus.Info(v)
	// }

	// 文件上传
	// 设置文件路径
	fileName := "test/config.yaml"
	// 上传文件夹对象,io流
	fi := strings.NewReader("test")

	// put提交请求
	if _, err := config.TxOss.Object.Put(context.Background(), fileName, fi, nil); err != nil {
		logrus.Info("put 预处理失败:", err)
	}
	// 上传本地文件
	if _, err := config.TxOss.Object.PutFromFile(context.Background(), fileName, "获取当前路径.go", nil); err != nil {
		logrus.Info("上传文件失败:", err)
	}

}
