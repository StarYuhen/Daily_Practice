package main

import (
	"fmt"
	"sync"
)

//这一文件主要是使用互斥锁

var must sync.Mutex
var mustR sync.RWMutex

/*
func main() {
	for  {
		go func() {
			must.Lock()
			fmt.Println("测试锁")
		}()
	}
}


*/
//这里可以发现一旦使用了look，将函数锁住，即使go在不断创建，可以依旧被锁起来了，无法执行下一步语句

/*
func main() {
	for i:=0;i<=100;i++ {
		go func() {
			//写锁定
			must.Lock()
			fmt.Println("这是被锁上的内容",i)
			if i%2==0{
				//写解除
				must.Unlock()
				fmt.Println("这是被解开的锁",i)
			}
		}()
	}
}

*/

//从这一段代码可以发现，倘若不解开锁，他只会执行着一遍的内容，剩下的全部都会堵塞。
/*
我有趣的发现当这样执行时(值的是不解锁)，i的值是会变的，他也是在一定时间内反应过来的。
执行锁定和解锁都需要花费时间，当然对于实际应用场面这点时间可以忽略不计。
*/

/*
func main() {
	for i:=0;i<=100;i++ {
		go func() {
		//读锁定
			mustR.RLock()
			fmt.Println("这是读锁定")
			//读解锁
			mustR.RUnlock()
		}()
	}
}

*/

//这么执行，你会发现go协程没有被堵塞。那是应为fmt包并没有写入，单纯的读取是不会被暂停的

//只执行一次的方法
func main() {
	s := &sync.Once{}
	for {
		s.Do(func() {
			fmt.Println("sdfsd")
		})
	}
}
