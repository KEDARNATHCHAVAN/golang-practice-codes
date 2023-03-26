package main

import "fmt"

const number int = 44 //declaration of constant variable

func main() {

	//explicit declarations : we have to declare the type of variable

	var str1 string = "this is first string"
	fmt.Println(str1)
	fmt.Printf("the type of the str1 is %T\n", str1)

	var integer int = 33 // default integer value is 0
	fmt.Println(integer)
	fmt.Printf("the type of 33 is %T\n", integer)

	//implicit declarations : compiler will automatically recogonise the type of variable
	// to omit the 'var' keyword, we can use ":=" operator, it only work for varibles inside functions

	var str2 = "this is second string"
	fmt.Println(str2)
	fmt.Printf("The type of str2 is also %T\n", str2)

	var num = 33
	fmt.Println(num)
	fmt.Printf("The type of num is %T\n", num)

	num2 := 34
	fmt.Println(num2)
}
