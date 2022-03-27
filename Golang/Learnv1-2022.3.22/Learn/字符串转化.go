package main

import (
	"fmt"
	"strings"
)

func main() {
	TestTitle()
	TestToTitle()
	TestToLower()
	TestToUpper()
}

func TestTitle() {
	fmt.Println(strings.Title("her royal highness"))
}

func TestToTitle() {
	fmt.Println(strings.ToTitle("louD noises"))
}

// func TestToTitleSpecial() {
// }

func TestToLower() {
	fmt.Println(strings.ToLower("Gopher"))
}

// func TestToLowerSpecial() {
//
// }

func TestToUpper() {
	fmt.Println(strings.ToUpper("Gopher"))
}

// func TestToUpperSpecial() {
//
