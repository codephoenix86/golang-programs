package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	nums := []int{10, 20, 30, 40, 50}
	for _, num := range nums {
		ch <- num
	}
	close(ch)
	for num := range ch {
		fmt.Println(num)
	}
}
