package main

import (
	"fmt"
)

// 2021/10/04 30年河东，30年河西，下次再来盘你

func main() {
	// fmt.Println(string(53)+string(70))
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w", 4))
}

/*
有一个密钥字符串 S ，只包含字母，数字以及 '-'（破折号）。其中， N 个 '-' 将字符串分成了 N+1 组。

给你一个数字 K，请你重新格式化字符串，使每个分组恰好包含 K 个字符。特别地，第一个分组包含的字符个数必须小于等于 K，但至少要包含 1 个字符。两个分组之间需要用 '-'（破折号）隔开，并且将所有的小写字母转换为大写字母。

给定非空字符串 S 和数字 K，按照上面描述的规则进行格式化。
*/

/*

index-> 0 value-> 53 string value-> 5
index-> 1 value-> 70 string value-> F
index-> 2 value-> 51 string value-> 3
index-> 3 value-> 90 string value-> Z
index-> 4 value-> 45 string value-> -
index-> 5 value-> 50 string value-> 2
index-> 6 value-> 69 string value-> E
index-> 7 value-> 45 string value-> -
index-> 8 value-> 57 string value-> 9
index-> 9 value-> 45 string value-> -
index-> 10 value-> 87 string value-> W
 题目： 5F3Z-2E-9-W k=4 => 5F3Z-2E9W
*/

/* 这题不能用hash表
func licenseKeyFormatting(s string, k int) string {
	// 使用字符串api将字符串中所有小写字母转化为大写字母
	str:=strings.ToTitle(s)
	hashmap:= map[int]int32{}
	J:=0
	for index,value:=range str{
	//	 fmt.Println("index->",index,"value->",value,"string value->",string(value))
		if value!=45{
			hashmap[index]=value
			J++
		}

	}

	fmt.Println(J%k)
	var strs string
	if J%k==0{
		for index,value:=range hashmap{
			 fmt.Println("index->",index,"value->",value,"string value->",string(value))
			 if index%4==0{
				 strs=strs+"-"
			 }
			 strs=strs+string(value)
		}
	}
	fmt.Println("strs->"+strs)
	fmt.Println(hashmap,J)
	return str
}

*/

func licenseKeyFormatting(s string, k int) string {

	return ""
}
