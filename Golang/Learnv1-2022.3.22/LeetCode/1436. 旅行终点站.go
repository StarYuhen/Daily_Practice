package main

import "fmt"

func main() {
	hsh := map[string]bool{}
	hsh["wdnms"] = true
	fmt.Println(hsh["wdnms"])
}

func destCity(paths [][]string) string {
	// 新建一个hash表 map
	hashtab := map[string]bool{}
	for _, value := range paths {
		// 将paths[0] 也就是cityAi全部储存在hash表中 因为cityAi必定是起始城市，所以不可以用它来判断终点
		hashtab[value[0]] = true
	}

	for _, values := range paths {
		// 接下来直接遍历数组后判断hash表中是否存在cityBi这个不同的终点，没有的相同就说明终点是他，有的话则直接返回空字符串
		if !hashtab[values[1]] {
			return values[1]
		}
	}
	return ""
}
