package main

import "github.com/sirupsen/logrus"

func main() {
	logrus.Info(len("测"))
	// qrcode.WriteFile("https://www.codesuger.com/", qrcode.Medium, 256, "./blog_qrcode.png")
}
