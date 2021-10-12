package main

import "fmt"

/*
使用异或运算，将所有值进行异或
异或运算，相异为真，相同为假，所以 a^a = 0 ;0^a = a
因为异或运算 满足交换律 a^b^a = a^a^b = b 所以数组经过异或运算，单独的值就剩下了
 思路
*/

func main() {
	fmt.Println(10 ^ 0)
}

func singleNumber(nums []int) int {
	var i int
	for _, value := range nums {
		i = value ^ i
	}

	return i
}
