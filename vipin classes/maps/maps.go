package main

import (
	"fmt"
)

func main() {
	var temp = make(map[string]int)
	temp["honey"] = 1
	temp["money"] = 2
	temp["boney"] = 3

	// u can use the delete function to delete the map
	fmt.Println(temp)
	delete(temp, "loney")

	//for checking if the key is present
	panda, present := temp["honey"]
	fmt.Println(panda)
	fmt.Println(present)
}
