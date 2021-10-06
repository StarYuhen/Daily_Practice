package main

import "fmt"

func main() {
	if li()() {
		fmt.Println("sdf")
	}
}

func li() (string, bool) {
	return "wdnmd", true
}
