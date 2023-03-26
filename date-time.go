package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("This code was compiled at", t)

	d := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Go was launched at", d)
	fmt.Printf("The type of time is %T\n", d)
}
