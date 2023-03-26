package main

import "fmt"

type Student struct {
	prn    int
	name   string
	branch string
}

func main() {
	stu := Student{47, "kedarnath", "it"}
	fmt.Println("The properties of student are", stu)
	stu.changename()
	stu.prn = 50
	fmt.Println(stu.prn)
}

// methods are just like member functions of classes in C++
func (s Student) changename() { //here 's' is the receiver and Student is the identifier for reference
	s.name = "viraj"
	fmt.Println("the changed name is", s.name)
}
