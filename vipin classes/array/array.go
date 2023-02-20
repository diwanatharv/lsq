package main

import "fmt"

func main() {
	//var arr [5]int = [5]int{1, 2, 3, 4, 5}
	//for i := 0; i < len(arr); i++ {
	//	println(arr[i])
	//}
	var arr1 = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr1)
	// golang itna bekar hai ki array ka size ko  custom input lene bhi nahi dekh sakta

	var arr2 = [2][2]int{{1, 1}, {0, 0}}
	fmt.Println(arr2)

	var n int
	fmt.Scanln(&n)

}
