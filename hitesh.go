package main

import (
	"fmt"
	"sort"
)

//outside thuis var:=is not allowed

// public : this equilavent naming the  global variable public

//user sei input lena hai
// fmt.Println("Enter Your First Name: ")

//     // var then variable name then variable type
//     var first string

//     // Taking input from user
//     fmt.Scanln(&first)
//     fmt.Println("Enter Second Last Name: ")
//     var second string
//     fmt.Scanln(&second)

//scanner:=bufio.NewScanner(os.Stdin);
//scanner.Scan()     this gives result in the string
//input:=scanner.text()
//     // Print function is used to
//     // display output in the same line
//     fmt.Print("Your Full Name is: ")

// // Addition of two string
// fmt.Print(first + " " + second)
func main() {
	// fmt.Println("hello");

	// var username string = "kamesh"
	// fmt.Printf("%v\n%T", username, username)

	// var isloggedin bool = false
	// fmt.Println(isloggedin)

	// panda := "handa"
	// fmt.Println(panda)

	//instead of try catch
	//input,ok:=
	//u can __
	//strconv used for converting string into other data types
	//strings library is always used
	// reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n')
	// numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(numrating + 1)
	// }

	//time package

	// presentTime := time.Now()
	// fmt.Println(presentTime)
	// //can get in a specific format
	// fmt.Println(presentTime.Format("01-02-2006 Monday"))

	//new  and make  used for dynamic memory allocation
	//make is intailized
	//pointer *after variable and & is also used for variable
	// var kamesh int = 9
	// var ptr *int = &kamesh
	// fmt.Println(ptr, *ptr)

	//array
	var fruitlist [4]string
	fruitlist[0] = "potato"
	fruitlist[1] = "apple"
	fruitlist[2] = "orange"
	fruitlist[3] = "kiwi"

	//agar size nahi pata aur bass dalna hai jo pata hai
	var arr3 = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr3)
	//slice
	//same as dynamic array
	//only difference from the normal array is there that if we copy some other array
	//it points to same memory location
	//cap() sei capacity nikal sakte ho
	/*var arr []int = []int{1, 2, 3, 4}*/
	//elements append karne kei liye

	//slicing can also be done in the array or the slice
	// var arr4 []int = []int{1, 2, 3, 4}
	// arr5 := append(arr4, 5, 6, 7, 8, 9, 10)
	// fmt.Println(arr5)
	//way of defining digitakky
	highscore := make([]int, 1)
	highscore[0] = 100
	highscore = append(highscore, 90, 100, 120, 130, 140, 150)
	panda := highscore
	fmt.Println(highscore)
	sort.Ints(highscore)
	fmt.Println(highscore)
	fmt.Println(panda)

	//maps
	empsal := make(map[string]int)
	empsal["kamesh"] = 27
	empsal["chait"] = 29
	empsal["ritik"] = 30
	fmt.Println(empsal)
	delete(empsal, "ritik")
	_, ok := empsal["chait"]
	fmt.Println(ok)

	//loops and ranges

	for key, value := range empsal {
		fmt.Printf("value is %+v %+v ", key, value)
	}
	//+v better way of

	day := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	for i := 0; i < len(day); i++ {
		fmt.Printf("%v\n ", day[i])
	}
	for i := range day {
		fmt.Printf("%v\n ", day[i])
	}
	s := proAdd(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(s)
}
func greet(name string) (bool, string) {
	fmt.Println("Hello" + name)
	return true, "greeting is done"
}
func proAdd(values ...int) int {
	total := 0

	for val := range values {
		total += val
	}
	return total
}
