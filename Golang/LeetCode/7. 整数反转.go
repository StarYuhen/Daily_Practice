package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(reverse(-2147483412))
}

func reverse(x int) int {
	xy, res, lens, lenst, lensts := 0, 0, len(strconv.Itoa(x)), 2147483647, -2147483648
	for i := 1; i <= lens; i++ {
		// 123 -> 321
		xy = x % 10
		x = x / 10
		if i == lens && xy == 0 {
			return res
		}
		res = res*10 + xy
		if res > lenst-1 || res < lensts {
			return 0
		}
	}
	return res
}
