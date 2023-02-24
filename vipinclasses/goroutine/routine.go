package main

import (
	"fmt"
	"time"
)

// go routines is used for thread creation
// go keyword
// time.sleep(time.second())
// like setimeout in js and used in the asynchronus feature
// this will  be running in the background
func msg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg, "=", i)
	}
}
func msg1(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg, "=", i)
	}
}
func main() {
	go msg("hi")
	go msg1("byi")
	time.Sleep(time.Second)
}
