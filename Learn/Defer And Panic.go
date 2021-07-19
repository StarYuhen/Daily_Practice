package main

import "fmt"

/*
//测试defer执行顺序
func main() {
	defer func() {
		fmt.Println("这是第一个defer")
	}()

	defer func() {
		fmt.Println("这是第二个defer")
	}()

	defer func() {
		fmt.Println("这是第三个defer")

		defer func() {fmt.Println("这是第三个defer里面嵌套的defer")}()

	}()

	defer func() {fmt.Println("这是第四个defer")}()
}

// defer的作用是最后执行，而defer的执行顺序又像是栈，先进后出的样子


*/
/*
func main() {
	defer func() {
		panic("这是第一个defer里的panic")
	}()
	defer func() {fmt.Println("这是第二个defer")}()
	//panic("这是第二个panic")
}
//由于我们知道defer是自动推迟到最后执行的,而panic不会，所以就先抛出错误


*/

//当注释掉panic时，就会执行defer的内容 ，所以打印内容才会是
/*
这是第二个defer
panic: 这是第一个defer里的panic
*/
func main() {
	defer func() {
		defer func() { fmt.Println("这是第一个defer") }()
		panic("这是第一个defer里的panic")
	}()
	defer func() { fmt.Println("这是第二个defer") }()
	//panic("这是第二个panic")
}

//这就可以看出第一个defer即使有panic，但他依旧先执行了defer，答应内容是
/*
这是第二个defer
这是第一个defer
panic: 这是第一个defer里的panic
*/
