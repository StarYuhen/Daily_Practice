package main

/*
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
*/

func main() {

}

func rotate(nums []int, k int) {
	lsit := make([]int, len(nums))
	for index, value := range nums {
		lsit[(index+k)%len(nums)] = value
	}
	copy(nums, lsit)
}
