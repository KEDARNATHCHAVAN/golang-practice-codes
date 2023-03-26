package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter your name:")
	fmt.Scan(&name)

	//1. no need of paranthesis
	//2. no need of break statement (it is by default)
	//3. if you want to continue further execution of program after matching a case,
	// 	 you can use "fallthrough" keyword after a case (just like continue in C++)

	switch name {
	case "kedarnath":
		fmt.Println("Its your first name !")
	case "prasad":
		fmt.Println("Its your middle name !")
	case "chavan":
		fmt.Println("Its your sirname !")
	default:
		fmt.Println("Its not your name :(")
	}
}
