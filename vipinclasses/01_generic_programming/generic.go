package main

import "fmt"

type number interface {
	int64 | float64
}

// for further optimization  u can create v as interface and include all those
//comparable will be same and 2nd one will be different 
func generic[k comparable, v number](m map[k]v) v {
	var s v
	for _, t := range m {
		s += t
	}
	return s
}
func main() {
	ints := make(map[string]int64)
	floats := make(map[string]float64)
	ints["first"] = 34
	ints["second"] = 19
	floats["first"] = 35.98
	floats["second"] = 17.89
	fmt.Println(generic(ints))
	fmt.Println(generic(floats))
}
