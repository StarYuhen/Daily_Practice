package main

import (
	"fmt"
)

func main() {
	var c = []int{1, 2, 2, 1}
	var b = []int{2, 2}
	fmt.Println(intersect(c, b))
}

func intersect(nums1 []int, nums2 []int) []int {
	// 将nums1全部存入hash表 s1,nums2)
	hashmap := map[int]int{}
	var list []int
	for _, value := range nums1 {
		hashmap[value]++
	}
	for _, values := range nums2 {
		if hashmap[values] > 0 {
			list = append(list, values)
			hashmap[values]--
		}
	}
	return list
}
