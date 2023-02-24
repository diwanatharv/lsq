package main

import (
	"awesomeProject/polymorphism"
)

// this is  only encapsulation
func main() {
	//var firstname string
	//var lastname string
	//var age int
	//fmt.Scanln(&firstname)
	//fmt.Scanln(&lastname)
	//fmt.Scanln(&age)
	//p := oops.PersonCreate(firstname, lastname, age)
	//fmt.Println(p.Getdetails())

	//this is for polymorphism
	// use the common interface and in the right hand side use any struct which is implementing all its
	//functions
	var c polymorphism.Shape = polymorphism.Circle{}
	var s polymorphism.Shape = polymorphism.Square{}
	c.Render()
	s.Render()
}
