package main

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func main() {
	// 初始化socket链接
	_, err := amqp.Dial("amqp://mrajoe:mrajoe@localhost:5672/mrajoe_vhost")
	logrus.Info(err)
}
