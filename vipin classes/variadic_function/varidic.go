package main

import "fmt"

// this is used for the when multiple parameters are passed
func sum(nums ...int) int {
	s := 0
	for _, val := range nums {
		s += val
	}
	return s
}

// for array
func sum1(nums []int) int {
	s := 0
	for _, val := range nums {
		s += val
	}
	return s
}
func main() {
	var arr []int = []int{1, 2, 3, 4, 9}
	fmt.Println(sum1(arr))
}
