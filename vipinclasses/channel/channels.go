package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getrandom() int {
	return rand.Intn(100)
}
func displaymessage(ch chan string) {
	for i := 1; i < 10; i++ {
		// time.Sleep(time.Second)// iske badla walthrough 
		ch <- fmt.Sprintf("this is value %d in channel", getrandom())
	}
	close(ch)
}
func main() {
	//used to make channel in golang pass
	// channel:=make(chan string )
	ch := make(chan string, 1)
	//special variables that is used to exchange imformation among goroutines
	//ulta c++ ka pointer is used to insert values
	//ch <- "this is value 1 in the channel"
	//ch <- "this is value 2 in the channel"
	//ch <- "this is value 3 in the channel"
	//fmt.Println(<-ch)
	// giving deadlock error as no buffer is there so we can pass the size by using the comma thing
	//u can print multiple values through range
	//close(ch)
	go displaymessage(ch)
	for chv := range ch {
		fmt.Println(chv)
	}

	//this range is also giving error
	//as it is trying to access other process
	//close function will shut the down the channel after there last message
	//used as  defer keyword
}
