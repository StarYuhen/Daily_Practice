package main

/*
给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。
*/

func main() {

}

func missingNumber(nums []int) (x int) {
	for i, v := range nums {
		x ^= i ^ v
	}
	return x ^ len(nums)
}
