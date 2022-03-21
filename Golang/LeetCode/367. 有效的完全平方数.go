package main

import (
	"github.com/sirupsen/logrus"
	"math"
)

/*
给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，否则返回 false 。

进阶：不要 使用任何内置的库函数，如  sqrt 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-perfect-square
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	logrus.Info(isPerfectSquare(14))
}

/*
平方根的定义：
完全平方指用一个整数乘以自己例如1*1，2*2，3*3等，依此类推。若一个数能表示成某个整数的平方的形式，则称这个数为完全平方数。完全平方数是非负数，
*/

func isPerfectSquare(num int) bool {
	// 当数字的平方根相乘等于原始数字就是完全平方根
	ints := int(math.Sqrt(float64(num)))
	return ints*ints == num
}
