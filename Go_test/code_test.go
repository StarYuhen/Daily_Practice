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



单纯这么调用另一个文件的参数可能造成函数未定义，我这里有用Goland的三个解决方案
1.将单元测试文件和包含的测试文件不放在同一目录，然后导用需要测试的文件所在目录即可
2.命令行打包的时候把文件一起打包  go test -v code.go  code_test.go
3.直接把要测试的哪个函数写在这个测试文件里面，然后直接测试即可

 这个案例使用第一种方法
*/

//普通的单元测试
func TestAdd(t *testing.T) {
	fmt.Println(code.Add(100))
}
