package main

func main() {
	s := []int{0, 1, 0, 3, 12}
	moveZeroes(s)
}

/*
func moveZeroes(nums []int)  {
	// 使用双指针
	j:=0
	for i:=0;i<len(nums);i++{
		if nums[i]==0{
			j++
		}else if j!=0 {
			nums[i-j]=nums[i]
			nums[i]=0
		}
	}
	fmt.Println(nums)
}

*/

func moveZeroes(nums []int) {
	// 双指针解法
	j := 0 // 累计当前零
	for i := 0; i < len(nums); i++ {
		// 如果现在是零就累计，执行完就开启下一次循环
		if nums[i] == 0 {
			j++
		} else if j != 0 {
			// 这里的语句意思是当nums[i]不等于零时，且出现过零，就会对出现零的数字做减法，将是0的下标变为当前不是0下标的数组值，然后将当前下标不是零的变成零
			nums[i-j] = nums[i]
			nums[i] = 0
		}
	}
}
