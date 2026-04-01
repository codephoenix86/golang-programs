package main

import "fmt"

func greet(name string, city string) string {
	return fmt.Sprintf("Hello, %s! Welcome from %s.", name, city)
}

func main() {
	name := "Naresh"
	city := "Nimbahera"
	fmt.Println(greet(name, city))
}
