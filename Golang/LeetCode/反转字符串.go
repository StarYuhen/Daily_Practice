package main

import "fmt"

/*

编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。

不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnhbqj/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func main() {
	reverseString([]byte("hello"))
}

func reverseString(s []byte) {
	var t []byte

	for i := len(s) - 1; i >= 0; i-- {
		t = append(t, s[i])
	}
	copy(s, t)
	fmt.Println("s-->", s, "t-->", t)
}
