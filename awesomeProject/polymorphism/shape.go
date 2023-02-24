package polymorphism

import "fmt"

type Shape interface {
	Render()
}

type Circle struct {
}

func (c Circle) Render() {
	fmt.Println("circle is getting renderered")
}

type Square struct {
}

func (c Square) Render() {
	fmt.Println("square is getting renderered")
}
