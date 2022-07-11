package main

import (
	"github.com/sirupsen/logrus"
	"strings"
)

func main() {
	ua := "Mozilla/5.0 (Linux; Android 7.0;) AppleWebKit/537.36 (KHTML, like Gecko) Mobile Safari/537.36 (compatible; PetalBot; https://webmaster.petalsearch.com/site/petalbot)"
	// 因为爬虫都有compatible; PetalBot;这种形式所以直接判断有没有compatible;就行
	test := strings.Contains(ua, "compatible;")
	logrus.Info(test)
}
