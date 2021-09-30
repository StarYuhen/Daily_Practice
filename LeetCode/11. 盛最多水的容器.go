package main

import "math"

func maxArea(height []int) int {
	// 双指针
	res, i, j := 0, 0, len(height)-1
	for i < j {
		// 计算
		ares := (j - i) * int(math.Min(float64(height[i]), float64(height[j])))
		res = int(math.Max(float64(res), float64(ares)))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}
