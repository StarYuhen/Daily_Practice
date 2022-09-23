package main

import "C"

func main() {
	// 利用Golang生成So供给Android调用
	// https://ejin66.github.io/2018/09/15/go-to-so-android.html
}

//export Hello
func Hello(message string) string {
	return "Hello," + message
}
