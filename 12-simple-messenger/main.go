package main

import (
	"fmt"
	"time"
)

func greet(message chan string) {
	time.Sleep(2 * time.Second)
	message <- "Hello from the goroutine!"
}

func main() {
	message := make(chan string)
	go greet(message)
	fmt.Println(<-message)
}
