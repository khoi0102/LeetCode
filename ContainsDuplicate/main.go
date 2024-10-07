package main

import "fmt"

func main() {
	fmt.Println("hello")
}
func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false

}
