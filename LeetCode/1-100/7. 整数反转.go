package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(reverse(-2147483648))
}

func reverse(x int) int {
	xy, res, lens, lenst, lensts := 0, 0, len(strconv.Itoa(x)), 33554431, -33554432
	for i := 1; i <= lens; i++ {
		// 123 -> 321
		if res > lenst-1 || res < lensts {
			return 0
		}
		xy = x % 10
		x = x / 10
		if i == lens && xy == 0 {
			return res
		}
		res = res*10 + xy
	}
	return res
}
