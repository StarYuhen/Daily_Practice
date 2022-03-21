package main

import "github.com/sirupsen/logrus"

/*
有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。



示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false


提示:

1 <= s.length, t.length <= 5 * 104
s 和 t 仅包含小写字母


作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn96us/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func main() {
	logrus.Info(isAnagram("a", "ab"))
}

func isAnagram(s string, t string) bool {
	hashmap := map[int32]int{}

	for _, v := range s {
		hashmap[v]++
	}

	for _, v := range t {
		if hashmap[v] != 0 {
			hashmap[v]--
		} else {
			hashmap[v]++
		}

	}

	for _, v := range hashmap {
		if v != 0 {
			return false
		}
	}
	return true
}
