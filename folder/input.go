package main

import "fmt"

func main() {
	var first, second string
	fmt.Print("Enter your first name: ")
	fmt.Scan(&first)
	fmt.Print("Enter your second name: ")
	fmt.Scan(&second)
	fmt.Println("Your name is ", first+" "+second)
}
