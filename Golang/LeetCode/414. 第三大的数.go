package main

import (
	"fmt"
	"sort"
)

func main() {
	var c = []int{1, 2, 2, 5, 3, 5}
	fmt.Println(thirdMax(c))
}

func thirdMax(nums []int) int {
	// 利用函数对数组进行降序排序
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	// 记录这是第几个数  需要返回第三大的数，如果没有，则返回最大的
	j := 0
	ma, min := nums[0], 0
	// 创建hash表记录重复的数字 bool默认的是false
	hashmap := map[int]bool{}
	for _, value := range nums {
		if !hashmap[value] {
			j++
			hashmap[value] = true
			min = value
		}

		if j == 3 {
			return value
		}

	}

	max := Max(ma, min)

	if j < 3 {
		return max
	}

	return max

}

func Max(ma int, min int) int {
	if ma > min {
		return ma
	} else {
		return min
	}
}
