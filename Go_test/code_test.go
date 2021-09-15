package Go_test

import (
	"./code"
	"fmt"
	"testing"
)

/*
利用testing这个包进行单元测试和压力测试  记住，测试的函数前面必须加TestXXX,如下面的TestAdd,关联的是另一个文件的Add函数
详细看这篇文章 https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter09/09.1.html
			https://learnku.com/docs/build-web-application-with-golang/how-113-go-writes-test-cases/3224
*/

//普通的单元测试
func TestAdd(t *testing.T) {
	fmt.Println(code.Add(100))

}
