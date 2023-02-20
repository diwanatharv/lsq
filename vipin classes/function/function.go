package main

import "fmt"

func add(val1 int, val2 int) int {
	return val1 + val2
}
func best(val1 int, val2 int) (int, int) {
	return val1 + val2, val1 - val2
}
func main() {
	fmt.Println(best(2, 5))
}
