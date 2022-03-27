package main

import "fmt"

/*
https://www.cnblogs.com/huansky/p/13488234.html

Sliding window algorithm is used to perform required operation on specific window size of given large buffer or array.

滑动窗口算法是在给定特定窗口大小的数组或字符串上执行要求的操作。

This technique shows how a nested for loop in few problems can be converted to single for loop and hence reducing the time complexity.

该技术可以将一部分问题中的嵌套循环转变为一个单循环，因此它可以减少时间复杂度。
*/

func main() {
	hashMap := map[string]int{}
	hashMap["string"] = 10
	fmt.Println(hashMap["String"])
	fmt.Println("最大连续字数:", lengthOfLongestSubstring("abcabca"))
}

func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
