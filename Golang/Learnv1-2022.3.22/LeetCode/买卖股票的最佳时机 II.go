package main

func main() {

}

/*
给定一个数组 prices ，其中 prices[i] 是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）

*/

// 他并没有规定数据次数，所以可以遍历的时候直接判断当前和上一个哪个大哪个小，如果现在的比上一个大就直接卖了
func maxProfit(prices []int) int {
	i := len(prices)
	if i == 0 {
		return 0
	}
	j := 0

	for p := 1; p < i; p++ {
		if prices[p-1] < prices[p] {
			j = prices[p] - prices[p-1] + j
		}
	}

	return j
}
