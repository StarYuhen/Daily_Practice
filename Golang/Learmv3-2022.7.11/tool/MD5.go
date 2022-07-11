package main

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/sirupsen/logrus"
)

func main() {
	d := []byte([]byte("StarYuhen"))
	m := md5.New()
	m.Write(d)
	logrus.Info(hex.EncodeToString(m.Sum(nil)))
}
