package main

import (
	"math"
)

func main() {

}

/*
给定一个初始元素全部为 0，大小为 m*n 的矩阵 M 以及在 M 上的一系列更新操作。

操作用二维数组表示，其中的每个操作用一个含有两个正整数 a 和 b 的数组表示，含义是将所有符合 0 <= i < a 以及 0 <= j < b 的元素 M[i][j] 的值都增加 1。

在执行给定的一系列操作后，你需要返回矩阵中含有最大整数的元素个数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/range-addition-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxCount(m int, n int, ops [][]int) int {
	// 先将数组进行顺序排序
	for _, v := range ops {
		m = int(math.Min(float64(m), float64(v[0])))
		n = int(math.Min(float64(n), float64(v[1])))
	}

	return m * n
}
