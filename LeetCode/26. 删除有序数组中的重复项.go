package main

import "fmt"

func main() {
	var c = []int{1, 1, 2}
	fmt.Println(removeDuplicates(c))
}

// 这道题我打算用hash表而不是双指针，哈希更方便也更块
// hash无法对应修改表数据，但可以统计数组中有哪些不重复
// 再来，我就是要用hash表
func removeDuplicates(nums []int) int {
	hashtap, j, i := map[int]bool{}, 0, len(nums)
	if i == 0 {
		return 0
	}
	for index, value := range nums {
		if hashtap[value] {
			continue
		}
		nums[j] = nums[index]
		j++
		hashtap[value] = true
	}
	return j
}

/*

// 老老实实用双指针 我就不信了，我非要用hash表 划掉，不能用hash表

func removeDuplicates(nums []int) int {
	i,j:=len(nums),1
	if i==0{
		return 0
	}
	for o:=1;o<i;o++{
		if nums[o]!=nums[o-1]{
			nums[j]=nums[o]
			j++
		}
	}
	fmt.Println(nums)
	return j
}


*/
