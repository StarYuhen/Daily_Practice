package main

import "fmt"

func main() {
	phoneNumber := "6283845388395"
	content := phoneNumber[2:len(phoneNumber)]
	fmt.Println(content)
}
