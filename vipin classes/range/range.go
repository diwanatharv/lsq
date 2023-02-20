package main

import "fmt"

func main() {

	//range is a loop
	//used for arr,slice and map
	var data []int = []int{1, 2, 3, 4}
	for index, value := range data {
		fmt.Println("index number is", index, "value is ", value)
	}
}
