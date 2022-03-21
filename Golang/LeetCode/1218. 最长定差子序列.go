package main

/*
给你一个整数数组 arr 和一个整数 difference，请你找出并返回 arr 中最长等差子序列的长度，该子序列中相邻元素之间的差等于 difference 。

子序列 是指在不改变其余元素顺序的情况下，通过删除一些元素或不删除任何元素而从 arr 派生出来的序列。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-arithmetic-subsequence-of-given-difference
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {

}

func longestSubsequence(arr []int, difference int) (ans int) {
	dm := map[int]int{}
	for _, v := range arr {
		dm[v] = dm[v-difference] + 1
		if dm[v] > ans {
			ans = dm[v]
		}
	}

	return ans
}
