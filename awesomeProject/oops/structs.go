package oops

type Person struct {
	firstName string
	lastName  string
	age       int
}

func (p Person) walk() string {
	return p.firstName + "is walking" + p.lastName
}

// constructor
func PersonCreate(f string, l string, a int) Person {
	// can add validation logic
	return Person{
		firstName: f,
		lastName:  l,
		age:       a,
	}
}
func (p *Person) Setdetails(s string, l string, a int) {
	p.firstName = s
	p.lastName = l
	p.age = a
}
func (p Person) Getdetails() (string, string, int) {
	return p.firstName, p.lastName, p.age
}
