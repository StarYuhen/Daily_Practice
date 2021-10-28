package main

import (
	"github.com/sirupsen/logrus"
	"math"
)

/*
字符串转换整数 (atoi)
请你来实现一个 myAtoi(string s) 函数，使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 atoi 函数）。

函数 myAtoi(string s) 的算法如下：

读入字符串并丢弃无用的前导空格
检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤 2 开始）。
如果整数数超过 32 位有符号整数范围 [−231,  231 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −231 的整数应该被固定为 −231 ，大于 231 − 1 的整数应该被固定为 231 − 1 。
返回整数作为最终结果。
注意：

本题中的空白字符只包括空格字符 ' ' 。
除前导空格或数字后的其余字符串外，请勿忽略 任何其他字符。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnoilh/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

/*
time="2021-10-28T20:51:03+08:00" level=info msg=49-----1
time="2021-10-28T20:51:03+08:00" level=info msg=50-----2
time="2021-10-28T20:51:03+08:00" level=info msg=51-----3
time="2021-10-28T20:51:03+08:00" level=info msg=52-----4
time="2021-10-28T20:51:03+08:00" level=info msg=53-----5
time="2021-10-28T20:51:03+08:00" level=info msg=54-----6
time="2021-10-28T20:51:03+08:00" level=info msg=55-----7
time="2021-10-28T20:51:03+08:00" level=info msg=56-----8
time="2021-10-28T20:51:03+08:00" level=info msg=57-----9
time="2021-10-28T20:51:03+08:00" level=info msg=48-----0
*/

func main() {
	logrus.Info(myAtoi(""))
	// Textnumber()
}

func myAtoi(s string) int {
	abs, sign, i, n := 0, 1, 0, len(s)
	// 丢弃无用的前导空格
	for i < n && s[i] == ' ' {
		i++
	}
	// 标记正负号
	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		}
	}
	for i < n && s[i] >= '0' && s[i] <= '9' {
		abs = 10*abs + int(s[i]-'0')  // 字节 byte '0' == 48
		if sign*abs < math.MinInt32 { // 整数超过 32 位有符号整数范围
			return math.MinInt32
		} else if sign*abs > math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}
	return sign * abs
}

func Textnumber() {
	s := "-+1234567890"
	for _, v := range s {
		logrus.Info(v, "-----", string(v))
	}
}
