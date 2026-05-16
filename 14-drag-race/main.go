package main

import (
	"fmt"
	"time"
)

func slowCar(message chan string) {
	time.Sleep(3 * time.Second)
	message <- "Slow car finished!"
}

func fastCar(message chan string) {
	time.Sleep(time.Second)
	message <- "Fast car finished!"
}

func main() {
	fastCh := make(chan string)
	slowCh := make(chan string)

	go slowCar(slowCh)
	go fastCar(fastCh)

	select {
	case message := <-fastCh:
		fmt.Println(message)
	case message := <-slowCh:
		fmt.Println(message)
	}
}
