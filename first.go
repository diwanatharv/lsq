// // remember one thing that this langauge is case sensitive
// // Golang is compiled
// // Used in cloud infrastructure
// // Lexer provides semicolon
// // //Go mod init hello
// // Package can be user defined
// // Import ftm is system based
// // Package and function name should be there
// // Println in new line and print in the same line
// // Package has to be always main
// // Default value is zero if we do not initialize a variable
// // If declared and not used it gives problem error
// // func main(){} ===     int main(){}
// // #include ===   import “fmt”
// // Package main
// // Print,printf,println
// // 3 types of data types basic type -num,string,boolean,aggregrate type-array,struct ,reference type-pointer,slice,map,function,channel ,interface type

// package main

// import "fmt"
// func test(s String, tyepc...)

// funString, tyepc...) Len() int           { return len(a) }
// funString, tyepc...) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// funString, tyepc...) Less(i, j int) bool { return a[i] < a[j] })  {

// }
// //after this has to be everything in new line

// // global level variable in (Pascal Case)
// // yei bass iss package mai accessible rahega
// //yaha par var hi vala decalaration
// /*var Val2 = 9*/

// //package level variable in camelCase
// // yei kisi bhi package mai accessible rahega
// /*var myVal = 10*/

// //enumerated const using iota
// //yei group mai apply hota hai
// //isme ek cheez aur jaise z ko define nahi kiye equal to iota kar kei phir vo kar lega
// //underscore sei skip hojata
// const (
// 	x = iota
// 	y = iota
// 	z = iota
// )

// //structure
// //same as c  used to make custom data type

// type Employee struct {
// 	Empid     int
// 	Empname   string
// 	deptname  string   //yei small case mai hai isliye different package mai nahi available ho raha hoga
// 	Empmobile []string //baaki sab available rahenege different package mai
// }

// type student struct {
// 	name        string
// 	phonenumber []string
// }

// func main() {
// 	// 1st type ka decalaration
// 	/*var a int
// 	a = 55*/
// 	// 2nd type ka decalaration
// 	/*var b int = 500*/

// 	//3rd type
// 	/*var c = 100*/

// 	//4th type
// 	/*d := 60
// 	fmt.Print("hello world ", a, " ", b, " ", c, " ", d)*/

// 	//scope of the variable -:
// 	//local  	var c = 100 //function kei andar jo bhi variables hote hai aur yahi par accessible hote hai
// 	//global
// 	//package level

// 	//shadowing
// 	//when local variable has same name as global .but local will be printed

// 	//printf
// 	//this is the same as c only a format specifier is used
// 	//%v for int
// 	//%T for type of the variable

// 	/*var number1 = 100.5
// 	fmt.Printf("%v ,%T", number1, number1)
// 	*/

// 	/*var a int32 = 234
// 	var b int64 = int64(a)
// 	fmt.Printf("%v %T ", a, b)
// 	*/

// 	/*fl := 7.6
// 	cl := float64(fl)
// 	fmt.Printf("%v %T ", fl, cl)
// 	*/

// 	//conversion of int to string
// 	/*var str int = 65
// 	var d string = strconv.Itoa(str)*/

// 	//conversion of ASCII Code
// 	/*var str int = 65
// 	var d string = string(str)
// 	fmt.Printf("%v %T", d, d)*/

// 	//complex number
// 	/*l1 := 4 + 2i
// 	fmt.Printf("%v %T", l1, l1)*/

// 	/*c := complex(2, 5)
// 	fmt.Printf("%v  %T", c, c)*/

// 	// can extract different thing from complex using real(),imag()

// 	//constant
// 	/*const i int = 5
// 	fmt.Print(i)*/

// 	//enumerated constant mai iota assign hota hai
// 	/*fmt.Println(x)
// 	fmt.Println(y)
// 	fmt.Println(z)
// 	*/

// 	//decalaring array
// 	/*var arr [4]int = [4]int{1, 2, 3, 4}*/

// 	//agar size nahi pata aur bass dalna hai jo pata hai
// 	/*var arr1 = [...]int{1, 2, 3, 4, 5, 6}*/
// 	/*fmt.Println(arr)*/
// 	/*fmt.Println(arr1)*/

// 	//len is used to print the lenght of the array

// 	/*fmt.Println(len(arr1))*/

// 	//if i want to change the index of some particular element
// 	// arr1[0] = 6
// 	/*fmt.Println(arr1)*/

// 	//copying it into entirely a new array

// 	/*arr2 := arr1
// 	fmt.Println(arr2)
// 	*/

// 	//only showing from a particular index to some other particular index

// 	/*fmt.Print(arr1[1:5])*/

// 	//matrix
// 	/*var matrix [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
// 	fmt.Println(matrix)*/

// 	//map
// 	//it stores key value pair

// 	//decalaration
// 	/*empsal := make(map[string]int)*/

// 	//intialization
// 	/*empsal = map[string]int{
// 		"Neha":   500,
// 		"kamesh": 980,
// 		"honey":  76,
// 	}
// 	*/
// 	//decalaration+intialization
// 	/*empsal := map[string]int{
// 		"Neha":   500,
// 		"kamesh": 980,
// 		"honey":  76,
// 	}
// 	empsal["honey"] = 98 */ //updating the previous one also a new can be added if there is different key

// 	//doubt is there any concept of ordered map like in case of c++

// 	//if u copy a map to other copy it will act in the same address location
// 	//that means it will change if u change of any of the contents in any of the map
// 	/*fmt.Println(empsal)*/

// 	//deleting a key

// 	/*delete(empsal, "honey")
// 	fmt.Println(empsal)*/

// 	//checking if a key is there in a map or not
// 	//here ok which is a boolean variable tells the bname can be anything
// 	//vaise ankit kei jagah underscore likhana jyaada accha rahega
// 	//go mai jo variable nahi pata uske badla underscore use kar sakte hi
// 	/*ankit, ok := empsal["kamesh"]
// 	fmt.Print(ankit, ok)*/

// 	//slice
// 	//same as dynamic array
// 	//only difference from the normal array is there that if we copy some other array
// 	//it points to same memory location
// 	//cap() sei capacity nikal sakte ho
// 	/*var arr []int = []int{1, 2, 3, 4}
// 	//elements append karne kei liye

// 	arr = append(arr, 1000)
// 	arr1 := append(arr, 98)
// 	//ekdum 2 array ko merge karne kei liye u have to use  ...
// 	arr2 := append(arr1, arr...)
// 	fmt.Println(arr2)

// 	//structure intialization

// 	emp1 := Employee{
// 		57,
// 		"kamesh",
// 		"siera",
// 		[]string{"7697236195", "9826198956"},
// 	}

// 	emp2 := Employee{
// 		99,
// 		"honey",
// 		"lapp",
// 		[]string{"9826198956"},
// 	}
// 	fmt.Print(emp1)
// 	fmt.Print(emp2)

// 	//STRUCTURE mai ek dusre structure mai copy karne sei vo dono same memory location mai point nahi karte
// 	// but if u use & while copying it points to same memory location
// 	// can access the elements of the objects of the array by dot operator
// 	fmt.Print(emp2.Empid)
// 	*/
// 	// loops
// 	//same hi hai pehle intialization phir conditional checking
// 	//phir increment ya decrement

// 	//for loop
// 	/*for i := 0; i <= 5; i++ {
// 		fmt.Printf("best %v \n", i)
// 	}*/

// 	//while loop;
// 	/*t := 1
// 	for t <= 5 {
// 		if t%2 == 0 {
// 			fmt.Printf("best %v \n", t)
// 		} else { //jaha if {} khatam ho rha h vaha else chaklu
// 			fmt.Printf("best %v \n", t)
// 		}
// 		t++
// 	}
// 	*/
// 	//switch
// 	// same like c some case and default will be there
// 	//(resolved)doubt is there is no case like break as if there will be all case will be executed
// 	//????????????????even continue keyword
// 	// break can be used
// 	// fallthrough keyword sei jo case select hua uske baad vala case bhi selected hoga
// 	// also using , u can use 1,3,4 in case for multiple things
// 	//??????no set thing like in the data structure
// 	/*i := 5
// 	switch i {
// 	case 1:
// 		fmt.Println(i)
// 	case 5:
// 		fmt.Println(i)
// 	default:
// 		fmt.Print("default case printed ", i)

// 	}
// 	*/

// 	//defer keyword in this when used is called at last
// 	//this is used in the database calling for closing the connection 1st

// 	/*
// 		defer fmt.Println("i")
// 		fmt.Println("love")
// 		fmt.Println("India")

// 	*/

// 	// pointer
// 	//same as c++
// 	/*var x int = 5
// 	var b *int = &x
// 	fmt.Print(*b)*/

// 	/*var s1 *student
// 	fmt.Println(s1)

// 	//??????????? new keyword is used for the memory allocation
// 	s1 = new(student)
// 	fmt.Println(s1)
// 	*/
// 	/*fmt.Println(showdata())
// 	fmt.Println(addition(9, 7))
// 	*/
// 	/*sum, sub := cals(9, 9)
// 	fmt.Println(sum)
// 	fmt.Print(sub)
// 	*/
// 	/*num1 := 40
// 	fmt.Println(pointerkause(&num1))
// 	*/

// 	var obj Vehicle = Bike{
// 		"HONDA",
// 		"black",
// 		7500.0,
// 	}
// 	obj.showdetails()
// }

// // function
// //used func keyword then the name of the function () for argument taking and then the return type
// //if u dont want to return anything then no daqta type to be mentioned
// ///... yei vale kei use mai range use kar sakte ho
// func showdata() string {
// 	return "hoi"
// }
// func addition(x int, y int) int {
// 	return x + y
// }

// //even can return many things in a function just has to mention the return type that many times
// //plus , in time of returning

// func cals(x int, y int) (int, int) {
// 	return (x + y), (x - y)
// }

// //used by call value and call by reference

// //????????did not understand the concept of the anonomous function
// func pointerkause(num *int) int {
// 	*num = 20
// 	return *num
// }

// //structure is group of variables
// //interface is the group of the functions
// //sab methods ko define karna pardega
// //structure kei help lete haio

// type Vehicle interface {
// 	showdetails()
// 	showname() string
// }
// type Bike struct {
// 	name  string
// 	color string
// 	price float64
// }

// func (bike Bike) showdetails() {
// 	fmt.Println(bike.color)
// 	fmt.Println(bike.name)
// 	fmt.Println(bike.price)
// }
// func (bike Bike) showname() string {
// 	return bike.name
// }


