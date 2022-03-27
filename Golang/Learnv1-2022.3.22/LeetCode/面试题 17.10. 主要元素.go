package main

import "fmt"

func main() {
	var c = []int{1, 1, 2}
	fmt.Println(majorityElement(c))
}

// 这道题可以使用hash表来计数
/*
数组中占比超过一半的元素称之为主要元素。给你一个 整数 数组，找出其中的主要元素。若没有，返回 -1 。请设计时间复杂度为 O(N) 、空间复杂度为 O(1) 的解决方案。
*/
func majorityElement(nums []int) int {
	hashmap := map[int]int{}
	lens := len(nums) / 2
	for _, value := range nums {
		// 判断是否存在该值，不存在的话默认值就是0，因为是int类型，不存在就写入并将value值写为1，存在就将值+1
		if hashmap[value] != 0 {
			hashmap[value]++
		} else {
			hashmap[value] = 1
		}
		// 当计数的数字大于一半时就直接返回那个数字
		if hashmap[value] > lens {
			return value
		}

	}

	// 都不存在就返回-1
	return -1
}
