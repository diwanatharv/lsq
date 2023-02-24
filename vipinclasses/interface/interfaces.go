package main

import "fmt"

// baaki iska theory nikalana pardega
// program samjh lo defer last mai use rehta hai close karne kei liye coonection

type calculator struct {
	num1 int
	num2 int
}
type calculation interface {
	add() int
	mul() int
}

func (c calculator) add() int {
	return c.num1 + c.num2
}
func (c calculator) mul() int {
	return c.num1 * c.num2
}

func usecalculation(cc calculation) {
	fmt.Println(cc.add())
	fmt.Println(cc.mul())
}
func main() {
	c := calculator{23, 23}
	usecalculation(c)
}
