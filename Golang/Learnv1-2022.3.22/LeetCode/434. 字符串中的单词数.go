package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSegments(""))
}

/*
index-> 0 value-> 72 value string-> H
index-> 1 value-> 101 value string-> e
index-> 2 value-> 108 value string-> l
index-> 3 value-> 108 value string-> l
index-> 4 value-> 111 value string-> o
index-> 5 value-> 44 value string-> ,
index-> 6 value-> 32 value string->
index-> 7 value-> 109 value string-> m
index-> 8 value-> 121 value string-> y
index-> 9 value-> 32 value string->
index-> 10 value-> 110 value string-> n
index-> 11 value-> 97 value string-> a
index-> 12 value-> 109 value string-> m
index-> 13 value-> 101 value string-> e
index-> 14 value-> 32 value string->
index-> 15 value-> 105 value string-> i
index-> 16 value-> 115 value string-> s
index-> 17 value-> 32 value string->
index-> 18 value-> 74 value string-> J
index-> 19 value-> 111 value string-> o
index-> 20 value-> 104 value string-> h
index-> 21 value-> 110 value string-> n
*/

func countSegments(s string) int {
	/*
		暂存思路
		if s == "" {
			return 0
		}
		j := 0
		for _, value := range s {
			if value == 32  {
				j++
				fmt.Println("value->", value)
			}
		}

		return j

	*/
	/*
		a:=strings.TrimSpace(s)
		b:=strings.Split(a," ")
		if b[0]==""{
			return 0
		}
		return len(b)

	*/

	i := 0
	for index, value := range s {
		/*
			解析：
			index==0且value的值不是空格才会加
			当index!=0时就会哦按段字符串当前下标的上一个是不是32 ，是的话就会同时判断当前循环字符串的值是不是空格，不是的话就会自加
			还是判断控格的那个套路，和我最初思路差不多
			可惜我当初完全不知道还可以这么干，原来加个（）就完事了
		*/
		if (index == 0 || s[index-1] == 32) && value != 32 {
			i++
		}
	}

	return i
}
