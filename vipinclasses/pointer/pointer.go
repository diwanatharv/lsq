package main

import "fmt"

func main() {
	num := 45
	var p *int
	p = &num

	fmt.Println(*p)
}
