package main

import "fmt"

func sumIntegers(data []any) int {
	sum := 0
	for _, element := range data {
		switch value := element.(type) {
		case int:
			sum += value
		}
	}
	return sum
}

func main() {
	mixedConfig := []any{"hello", 42, true, 8, 3.14, "world", 10}
	result := sumIntegers((mixedConfig))
	fmt.Printf("Total sum of extracted integers: %d\n", result)
}
