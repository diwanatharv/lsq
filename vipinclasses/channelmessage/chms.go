package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getrand() int {
	return rand.Intn(100)
}
func getsum(ch chan string) {
	n1 := getrand()
	n2 := getrand()
	sum := n1 + n2
	ch <- fmt.Sprintf("the sum of 2 %d and %d random number is %d", n1, n2, sum)
	//yei vala line value insert kar raha hai
	time.Sleep(time.Millisecond * 500)
}
func getsum1(ch chan string) {
	fmt.Println(<-ch)
	//yei vala line value print kar raha hau
	time.Sleep(time.Millisecond * 500)
}
func main() {
	result := make(chan string)
	for i := 1; i < 10; i++ {
		time.Sleep(time.Second)
		go getsum(result)
		go getsum1(result)
	}
	close(result)
}
