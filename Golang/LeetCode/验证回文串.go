package main

import (
	"github.com/sirupsen/logrus"
	"log"
	"regexp"
	"strings"
)

/*
验证回文串
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。


示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
解释："amanaplanacanalpanama" 是回文串
示例 2:

输入: "race a car"
输出: false
解释："raceacar" 不是回文串

回文串=====>指的是无论从左边还是右边读结果都是一样的

提示：

1 <= s.length <= 2 * 105
字符串 s 由 ASCII 字符组成

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xne8id/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func main() {
	logrus.Info(isPalindrome(""))
}

func isPalindrome(s string) bool {
	// 将字符串全部转化为小写
	str := strings.ToLower(s)
	// 去除不是字母和数字的字符，使用，在字母和数字的范围是[^a-zA-Z0-9 ]
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	str = reg.ReplaceAllString(str, "")

	ls := []byte(str)
	lens := len(ls)
	lens--
	for _, value := range ls {
		if value != ls[lens] {
			return false
		}
		lens--
		// logrus.Info(value,"====",string(value))
	}

	return true
}
