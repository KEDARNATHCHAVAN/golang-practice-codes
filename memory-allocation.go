package main

import (
	"fmt"
)

func main() {

	//memory allocation

	m := make(map[string]int)
	m["key"] = 33
	fmt.Println(m)

	//pointeres

	a := 34.54
	ptr := &a
	fmt.Println("The value of pointer is", ptr)
	fmt.Println("The value at pointer is", *ptr)

	*ptr = *ptr + 1
	fmt.Println("The changed value at pointer is", *ptr)
}
