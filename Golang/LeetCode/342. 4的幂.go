package main

func main() {

}

func isPowerOfFour(n int) bool {
	return n > 0 && n&(n-1) == 0 && n%3 == 1
}
