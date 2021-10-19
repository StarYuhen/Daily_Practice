package main

import "fmt"

/*
字符串中的第一个唯一字符
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
示例：

s = "leetcode"
返回 0

s = "loveleetcode"
返回 2

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn5z8r/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

题目分析：将字符串全部写入字典，然后输出第一个字典为1的数,需要数组一个循环来返回当前在那里的索引，暴力解法 Goint类型默认值是0，byte好像是0x0
// 明天优化.jpg
*/

func main() {
	fmt.Println(firstUniqChar("aabb"))
}

func firstUniqChar(s string) int {
	hashmap := map[string]int{}
	// 循环遍历字符串内容并写入哈希表，然后计数
	for _, v := range s {
		hashmap[string(v)]++
	}

	for i := 0; i < len(s); i++ {
		if hashmap[string(s[i])] == 1 {
			return i
		}
	}

	return -1
}
