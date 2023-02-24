package main

import (
	"fmt"
)

type Person struct {
	firstname string
	lastname  string
	age       int
}

func CreatePerson(ft string, lt string, age int) Person {
	return Person{
		firstname: ft,
		lastname:  lt,
		age:       age,
	}
}
func Setdetails(s string, l string, a int, p *Person) {
	p.firstname = s
	p.lastname = l
	p.age = a

}

// func CreatePerson(ft string, lt string, age int,Person p) Person {
//
// }
func main() {
	//p1 := Person{
	//	Firstname: "kamesh",
	//  Lastname:  "verma",
	//	Age:       21,
	//}
	//
	//var ptr *Person = &p1
	//fmt.Println(ptr.Age)
	var firstname string
	fmt.Scanln(&firstname)
	var lastname string
	fmt.Scanln(&lastname)
	var age int
	fmt.Scanln(&age)
	//p := CreatePerson(firstname, lastname, age)
	p := Person{}
	Setdetails(firstname, lastname, age, &p)
	fmt.Println(p)
}
