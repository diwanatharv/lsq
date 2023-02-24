package main

import "fmt"

func main() {

	// concept of switch as if in the normal switch no comparison is allowed
	//eg i>5 else
	//so the  solution can be to after switch keyword nothing to switch should be given
	//case can be done than  compared mode
	//this is the normal switch case
	i := 5
	switch i {
	case 1:
		fmt.Println("i is this", i)
	case 5:
		fmt.Println("i is this", i)
	default:
		fmt.Println("i is this", i)
	}
}
