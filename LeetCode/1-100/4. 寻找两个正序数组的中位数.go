package main

import (
	"fmt"
	"math"
	"sort"
)

/*
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
*/

func main() {
	fmt.Println(float64(5 / 2))
	var name = []int{1, 2}
	var names = []int{3, 4}
	fmt.Println(findMedianSortedArrays(name, names))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 合并两个数组
	nums := append(nums1, nums2...)
	// 使用Go api将数组结合排序
	sort.Ints(nums)
	fmt.Println(len(nums))
	lens := len(nums)
	l := int(math.Floor(float64(lens / 2)))
	fmt.Println(l)
	fmt.Println(lens % 2)
	if lens%2 == 0 {
		// 说明这个数组是偶数个数 值就是中间两个数想和除以二得除数
		print("ws", (nums[l]+nums[l-1])/2)
		return (float64(nums[l]) + float64(nums[l-1])) / 2
	}

	return float64(nums[l])

}
