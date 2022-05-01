package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

func main() {
	m := gomail.NewMessage()

	// 发送人
	m.SetHeader("From", "sDFGDFG")
	// 接收人
	m.SetHeader("To", "3446623843@qq.com")
	// 抄送人
	// m.SetAddressHeader("Cc", "xxx@qq.com", "xiaozhujiao")
	// 主题
	m.SetHeader("Subject", "AliiothForum论坛发送的验证码")
	// 内容

	body := fmt.Sprintf("<a href=%s style=color:#12addb target=_blank>%s </a> <h2>验证码：%s<h2> <p>若此邮件不是您请求的，请忽略并删除！</p> ", "https://www.yuhenm.com", "AliiothForum论坛", "456457567")
	m.SetBody("text/html", body)
	// m.Attach("./myIpPic.png")

	// 拿到token，并进行连接,第4个参数是填授权码
	d := gomail.NewDialer("smtp.qq.com", 587, "3446623843@qq.com", "xsiujvorfegzcjfd")

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("DialAndSend err %v:", err)
		panic(err)
	}
	fmt.Printf("send mail success\n")
}
