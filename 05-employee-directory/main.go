package main

import "fmt"

type Employee struct {
	id    int
	name  string
	title string
}

func Promote(employee *Employee, target string) {
	employee.title = target
}

func main() {
	coder := Employee{
		id:    101,
		name:  "naresh",
		title: "Junior Developer",
	}

	fmt.Printf("Before: %+v\n", coder)
	Promote(&coder, "Senior Developer")
	fmt.Printf("After: %+v\n", coder)
}
