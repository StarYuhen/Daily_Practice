package main

import (
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

type Captcha struct {
	KeyLong   int `mapstructure:"key-long" json:"keyLong" yaml:"key-long"`       // 验证码长度
	ImgWidth  int `mapstructure:"img-width" json:"imgWidth" yaml:"img-width"`    // 图片宽度
	ImgHeight int `mapstructure:"img-height" json:"imgHeight" yaml:"img-height"` // 图片高度
}

var store = base64Captcha.DefaultMemStore

var capt = CaptchaConfig()

func CaptchaConfig() Captcha {
	// 获取验证码配置
	var config Captcha
	config.KeyLong = 6
	config.ImgWidth = 240
	config.ImgHeight = 80
	return config
}
func main() {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(capt.ImgHeight, capt.ImgWidth, capt.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		logrus.Info("验证码获取失败")
	} else {
		/*
			response.OkWithDetailed(response.SysCaptchaResponse{
				CaptchaId: id,
				PicPath:   b64s,
			}, "验证码获取成功", c)

		*/
		number := store.Get(id, true)
		s := store.Verify(id, number, true)
		logrus.Info("验证码信息--->", id, b64s, "------", number)
		logrus.Info("检验是否整确-->", s, "----")
	}
}
