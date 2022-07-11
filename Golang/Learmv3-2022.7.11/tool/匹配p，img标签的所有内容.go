package main

import (
	"github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

func main() {
	str := "<h2>首页</h2><p>手动阀手动阀撒阿道夫</p><p> sdfgdgGRDFTHRT更好的法国</p>"
	// (?m)<p.*?[^<]>.*?</p>
	var hrefRegexp = regexp.MustCompile("<p.*?>.*?</p>")
	match := hrefRegexp.FindAllString(str, -1)
	var content string
	if match != nil {
		for _, v := range match {
			// p标签提取
			list := strings.Split(strings.Split(v, ">")[1], "<")[0]
			content = content + list
			if len(content) > 550 {
				// 截取前50个
				logrus.Info(content)
				logrus.Info(content[:550])
				return
			}
		}
		logrus.Info(content)
	}

}
