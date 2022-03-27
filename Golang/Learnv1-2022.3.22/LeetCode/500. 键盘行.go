package main

import (
	"log"
	"strings"
)

/*
给你一个字符串数组 words ，只返回可以使用在 美式键盘 同一行的字母打印出来的单词。键盘如下图所示。

美式键盘 中：

第一行由字符 "qwertyuiop" 组成。
第二行由字符 "asdfghjkl" 组成。
第三行由字符 "zxcvbnm" 组成。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/keyboard-row
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	log.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}

// 使用标准库定义的方法
func findWords(words []string) []string {
	// 定义键盘的三行
	line1 := "q & w & e & r & t & y & u & i & o & p & Q & W & E & R & T & Y & U & I & O & P"
	line2 := "a & s & d & f & g & h & j & k & l & A & S & D & F & G & H & J & K & L"
	line3 := "z & x & c & v & b & n & m & Z & X & C & V & B & N & M"
	var res []string
	for _, v := range words {
		bool1 := strings.ContainsAny(v, line1)
		bool2 := strings.ContainsAny(v, line2)
		bool3 := strings.ContainsAny(v, line3)
		if (bool1 && !bool2 && !bool3) || (bool2 && !bool1 && !bool3) || (bool3 && !bool2 && !bool1) {
			res = append(res, v)
		}
	}

	return res

}
