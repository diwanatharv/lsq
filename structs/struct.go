package structs

type Person struct {
	firstName string //pRIVATE field
	lastName  string
	age       int
	//FullName  func() string  functions can be used in the structure
}

// reciever function
func (p Person) walk() string {
	return p.firstName + p.lastName + "Likes Walking"
}
