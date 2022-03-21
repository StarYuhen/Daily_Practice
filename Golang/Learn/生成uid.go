package main

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"github.com/sirupsen/logrus"
)

func main() {
	str, _ := uuid.NewV4()
	s := fmt.Sprintf("%s", str)
	logrus.Info(s)
}
