//package main
//
//import (
//	"errors"
//	"fmt"
//)
//
//// error and exception handling in
//// return extra field called error
//
//// can also create your custom type of error handling
//
//// whenever we have error in  the program we can do the panic  and it will not allow to move forward
//
////main idea of the recover is to recover from the panic function used above panic..maybe anonymous functionn
////is used with the defer  keyword
//type dividebyzero struct {
//	texterror string
//}
//
//func (c dividebyzero) Error() string {
//	return c.texterror
//}
//func dividedbyzero() error {
//	return dividedbyzero("your error is ")
//}
//func divison(num1 int, num2 int) (int, error) {
//	if num2 == 0 {
//		return -1, errors.New("division by zero")
//		//can add the custom type also
//	}
//	return num1 / num2, nil
//}
//func main() {
//	p, ok := divison(20, 0)
//	if ok != nil {
//		fmt.Println(ok.Error())
//	} else {
//		fmt.Println(p, ok)
//	}
//}
