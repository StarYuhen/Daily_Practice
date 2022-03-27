package main

import "fmt"

func main() {
	var s = []int{1, 2, 3}
	fmt.Println(plusOne(s))
}

/*
func plusOne(digits []int) []int {
	i:=len(digits)
	j:=digits[i-1]
	if j==9{
		digits[i-1]=1
		digits=append(digits,0)
	}else {
		digits[i-1]=j+1
	}
	return digits
}

*/

func plusOne(digits []int) []int {
	length := len(digits) - 1
	for i := length; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}

	digits = append([]int{1}, digits...)
	return digits
}
