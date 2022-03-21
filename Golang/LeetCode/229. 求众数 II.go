package main

import "github.com/sirupsen/logrus"

/*
给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
*/
func main() {
	var s = []int{3, 2, 3}
	logrus.Info(_majorityElement(s))
}

func _majorityElement(nums []int) []int {
	hashmap := map[int]int{}
	lens := len(nums) / 3
	var num []int
	for _, value := range nums {
		// 循环将数据写入
		hashmap[value]++
	}
	for index, value := range hashmap {
		if value > lens {
			num = append(num, index)
		}
	}
	return num
}
