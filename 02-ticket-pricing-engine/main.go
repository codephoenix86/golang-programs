package main

import "fmt"

func CalculateTicketPrice(age int, week string) int {
	price := 0

	switch week {
	case "Saturday", "Sunday":
		price += 2
	}

	if age < 13 {
		price += 8
	} else if age < 65 {
		price += 12
	} else {
		price += 10
	}

	return price
}

func main() {
	ages := []int{10, 25, 70, 12, 45}
	days := []string{"Monday", "Saturday", "Wednesday", "Sunday", "Friday"}

	for i := 0; i < 5; i++ {
		fmt.Printf("price = %d\n", CalculateTicketPrice(ages[i], days[i]))
	}
}
