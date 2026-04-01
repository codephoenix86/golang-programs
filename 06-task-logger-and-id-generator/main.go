package main

import "fmt"

func makeIDGenerator() func() int {
	id := 0

	generateId := func() int {
		id++
		return id
	}

	return generateId
}

func main() {
	defer fmt.Println("System shutdown complete.")

	generateId := makeIDGenerator()

	for i := 0; i < 5; i++ {
		fmt.Printf("Starting task %d...\n", generateId())
	}
}
