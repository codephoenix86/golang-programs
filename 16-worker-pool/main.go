package main

import (
	"fmt"
	"time"
)

func worker(jobs, results chan int) {
	for job := range jobs {
		job *= 2
		time.Sleep(500 * time.Millisecond)
		results <- job
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		jobs <- i
	}
	for range 3 {
		go worker(jobs, results)
	}
	close(jobs)
	for range 10 {
		fmt.Println(<-results)
	}
}
