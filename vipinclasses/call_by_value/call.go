package main

import "fmt"

func swap(num1 *int, num2 *int) {
	temp := *num1
	*num1 = *num2
	*num2 = temp

}
func main() {
	num1 := 3
	num2 := 5
	fmt.Println(num1, " ", num2)
	swap(&num1, &num2)
	fmt.Println(num1, " ", num2)
}
