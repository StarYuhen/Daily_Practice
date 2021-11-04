package main

/*
解题思路： https://www.yuhenm.com/index.php/archives/271/
*/
func main() {

}

func twoSum(nums []int, target int) []int {
	// 创建一个hashmap表
	hashMap := map[int]int{}
	for index, value := range nums {
		if p, ok := hashMap[target-value]; ok {
			return []int{p, index}
		}
		hashMap[value] = index
	}
	return nil
}

func twosumInit(nums []int, target int) []int {
	// 创建一个hashmap表 用于储存过滤过的数据
	hashMap := map[int]int{}
	// 利用循环来判断哈希表中是否存在满足条件的值 index下标 value值
	for index, value := range nums {
		// 这个判断是读取当target-value剩下的数有相同的，就说明当前value下标与储存在map中的当前下标是这道题的答案
		if i, ok := hashMap[target-value]; ok {
			return []int{i, index}
		}
		// 当条件不满足时就把值存入hash表，这样下一次循环就可以利用hash表的特性方便的查找是否有符合条件的值
		hashMap[value] = index
	}
	// 当条件都不满足时返回空
	return nil
}
