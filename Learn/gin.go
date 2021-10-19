package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	res := gin.Default()
	res.GET("/api", func(context *gin.Context) {
		logrus.Info(context.Query("uid"))
	})

	err := res.Run(":80")
	if err != nil {
		return
	}
}
