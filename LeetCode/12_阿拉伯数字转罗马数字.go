package main

import "fmt"

func main() {
	tests()
}

/*
	字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/integer-to-roman
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func intToRoman(num int) string {
	var (
		thousands = []string{"", "M", "MM", "MMM"}
		hundreds  = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
		tens      = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
		ones      = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	)
	return thousands[num/1000] + hundreds[num%1000/100] + tens[num%100/10] + ones[num%10]
}

func tests() {
	fmt.Println(2000 / 1000)
}
