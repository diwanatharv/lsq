package main

import "fmt"

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
func main() {
	//slices
	// in slices u do not need to be define the size (main thing )which differentiates it from array
	//dynamic array
	//u can intialize or not intialize it too
	// main thing that u can always add elemet from append function
	// u can remove function from  the slice using the slicing method

	var arr1 []int = []int{1, 2, 3, 4}
	arr1 = append(arr1, 5, 6, 7)
	//fmt.Println(arr1)

	//slicing
	fmt.Println(arr1[1:3]) // will start from  that index to n-1 index
	//gives out of bound of acces agar masti kiye toh

	//u can create slices using make keyword also and make keyword also use to make map
	// new keyword is used to create struct
	var n int
	fmt.Scanln(&n)
	var makeslice = make([]int, n)
	makeslice = append(makeslice, 1, 2)
	fmt.Println(makeslice)
}
