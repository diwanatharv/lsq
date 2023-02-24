package Student

import "fmt"

type Student struct {
	Name   string
	Course string
	Rollno int
}

func (s *Student) getstudent() {
	fmt.Printf("name of the student is %s , course is %s , rollno is %d", s.Name, s.Course, s.Rollno)
}
