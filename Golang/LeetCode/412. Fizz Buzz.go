package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzBuzz(3))
}

func fizzBuzz(n int) []string {
	var s []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			s = append(s, "FizzBuzz")
			continue
		}

		if i%3 == 0 {
			s = append(s, "Fizz")
			continue
		}

		if i%5 == 0 {
			s = append(s, "Buzz")
			continue
		}

		s = append(s, strconv.Itoa(i))

	}

	return s
}
