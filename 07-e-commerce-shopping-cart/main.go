package main

import "fmt"

type Cart struct {
	items []float64
}

func (cart *Cart) AddItem(price float64) {
	cart.items = append(cart.items, price)
}

func (cart *Cart) Total() float64 {
	total := 0.0

	for _, item := range cart.items {
		total += item
	}

	return total
}

func main() {
	cart := Cart{}

	cart.AddItem(19.99)
	cart.AddItem(5.50)
	cart.AddItem(12.00)

	fmt.Printf("Cart Total: %.2f\n", cart.Total())
}
