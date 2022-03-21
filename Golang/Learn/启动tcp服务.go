package main

import (
	"github.com/sirupsen/logrus"
	"net"
)

func main() {
	client, err := net.Dial("tcp", ":80")
	if err != nil {
		logrus.Info("error in ", err)
	}
	logrus.Info(client)
}
