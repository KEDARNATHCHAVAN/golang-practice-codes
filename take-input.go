package main

import (
	"bufio" //(buffered IO) temporarily accumulate the results for an IO operation before transmitting it forward
	"fmt"
	"os"
	"strconv" //for string conversion
	"strings" //contains functions for manipulating strings
)

func main() {
	read := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text : ")
	input, _ := read.ReadString('\n')
	fmt.Println("you have entered : ", input)

	//to convert a string into different type
	fmt.Print("Enter a number: ")
	num, _ := read.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(num), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number: ", aFloat)
	}
}
