package main

import (
	"fmt"
)

type Student struct {
	Name   string
	Roll   int
	Course string
}

func (s *Student) st() string {
	return s.Name + s.Course
}

type number interface {
	int64 | float64
}

// in generic programming there is k comparable and the other one is different number
func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	//sum := 0
	//for i := 1; i < 10; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)
	//range   array,slice ,map or channel in the string is used to iterate

	//s := Student{
	//	Name:   "Atharv",
	//	Roll:   9,
	//	Course: "IT",
	//}
	//fmt.Println(s)
	//fmt.Println(s.st())
	//ints := make(map[string]int)
	//ints["kamesh"] = 1
	//floats := make(map[string]int)
	//floats["santosh"] = 2
	//arr5 := [5]int{1, 2, 3, 4}
	//fmt.Println(arr5)
	//arr := [...]int{1, 2, 43, 98}
	//fmt.Println(arr)
	a := 5
	b := 6
	swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
	arr := []int{}
	arr = append(arr, 1, 9, 6)
	//multiple return is possible
	//Interfaces are used in Go where polymorphism is needed.
	//In a function where multiple types can be passed an interface can be used.
	//Interfaces allow Go to have polymorphism.
	//Interfaces are also used to help reduce duplication/boilerplate code.
	//
	//Interfaces are very useful in case of functions and methods where
	//you need argument of dynamic types, like Println function which accepts any type of values.
	//
	//When multiple types implement the same interface,
	//it becomes easy to work with them. Hence whenever we can use interfaces, we should.

	//used to make channel in golang pass
	// channel:=make(chan string )
	//ch := make(chan string, 1)
	////special variables that is used to exchange imformation among goroutines
	////ulta c++ ka pointer is used to insert values
	////ch <- "this is value 1 in the channel"
	////ch <- "this is value 2 in the channel"
	////ch <- "this is value 3 in the channel"
	////fmt.Println(<-ch)
	//// giving deadlock error as no buffer is there so we can pass the size by using the comma thing
	////u can print multiple values through range
	////close(ch)
	//go displaymessage(ch)
	//for chv := range ch {
	//	fmt.Println(chv)
	//}

	//// go routines is used for thread creation
	//// go keyword
	//// time.sleep(time.second())
	//// like setimeout in js and used in the asynchronus feature
	//// this will  be running in the background
	//func msg(msg string) {
	//	for i := 0; i < 5; i++ {
	//		fmt.Println(msg, "=", i)
	//	}
	//}
	//func msg1(msg string) {
	//	for i := 0; i < 5; i++ {
	//		fmt.Println(msg, "=", i)
	//	}
	//}
	//func main() {
	//	go msg("hi")
	//	go msg1("byi")
	//	time.Sleep(time.Second)
	//}

	// channels are the way through which multiple go routines interact
	//func main() {
	//	mych := make(chan int, 2)
	//	// pushing the value in the channel
	//	//mych <- 5
	//	//getting value from the channel
	//	//fmt.Println(<-mych)
	//	//if someone is listening channel then only it allows to pass the value
	//	wg := &sync.WaitGroup{}
	//	wg.Add(2)
	//	//1st function will recieve values
	//	//send only
	//	go func(ch <-chan int, wg *sync.WaitGroup) {
	//		fmt.Println(<-mych)
	//
	//		wg.Done()
	//	}(mych, wg)
	//
	//	// 2nd function will add values
	//	//recieve only
	//	go func(ch chan<- int, wg *sync.WaitGroup) {
	//		mych <- 5
	//		mych <- 6
	//		close(mych)
	//		wg.Done()
	//	}(mych, wg)
	//
	//	wg.Wait()
	//}

	//var score = []int{0}
	//wg := &sync.WaitGroup{}
	//mut := &sync.Mutex{}
	//wg.Add(3) //no .of go routines 3
	//go func(wg *sync.WaitGroup, mut *sync.Mutex) {
	//	fmt.Println("1st routine")
	//	mut.Lock()
	//	score = append(score, 1)
	//	mut.Unlock()
	//	wg.Done()
	//}(wg, mut)
	//go func(wg *sync.WaitGroup, mut *sync.Mutex) {
	//	fmt.Println("2nd routine")
	//	mut.Lock()
	//	score = append(score, 2)
	//	mut.Unlock()
	//	wg.Done()
	//}(wg, mut)
	//go func(wg *sync.WaitGroup, mut *sync.Mutex) {
	//	fmt.Println("3nd routine")
	//	score = append(score, 3)
	//	wg.Done()
	//}(wg, mut)
	//wg.Wait()
	//fmt.Println(score)
}
