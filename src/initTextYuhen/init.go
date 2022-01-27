package main

import "github.com/sirupsen/logrus"

// 完成JavaScript的原型链功能

// InterfaceTo
// 冷知识：创建的类型无论是不是interface都可以创建函数和方法，所以type创建接口和类型可以换一种方法理解：
// type创建一个储存方法的空间，而后面的类型则是方法指定的储存内容类型，并非必须需要interface。
// 提出疑问：如何才能已经套了一个函数的时候再用一个函数如：i.ToString().Msg() 实现方法如下
type InterfaceTo []byte

type InterfaceToV7 []byte

func (i *InterfaceTo) ToString() string {
	return "name"
}

func (i *InterfaceTo) ToStringIn() *InterfaceToV7 {
	return (*InterfaceToV7)(i)
}

func (v *InterfaceToV7) Tostring() {
	logrus.Infoln("测试成功")
}

func main() {
	logrus.Info(true)
	var i InterfaceTo
	// logrus.Infoln(i.ToString())
	// 实现方法成功
	logrus.Info(i.ToStringIn().Tostring)
}
