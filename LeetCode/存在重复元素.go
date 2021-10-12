package main

func main() {

}

func containsDuplicate(nums []int) bool {
	hashmap := map[int]bool{}
	for _, value := range nums {
		if hashmap[value] {
			return true
		}
		hashmap[value] = true

	}
	return false
}
