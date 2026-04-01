package main

import "fmt"

func GetTemperatureStats(temperatures []float64) (float64, float64, float64) {
	n := len(temperatures)

	if n == 0 {
		return 0, 0, 0
	}

	minTemp := temperatures[0]
	maxTemp := temperatures[0]
	total := 0.0

	for _, temp := range temperatures {
		if temp < minTemp {
			minTemp = temp
		}
		if temp > maxTemp {
			maxTemp = temp
		}
		total += temp
	}

	averageTemp := total / float64(n)

	return maxTemp, minTemp, averageTemp
}

func main() {
	temperatures := []float64{72.5, 75.0, 70.1, 80.2, 78.5}

	maxTemp, minTemp, averageTemp := GetTemperatureStats(temperatures)

	fmt.Printf("Min: %.2f\n", minTemp)
	fmt.Printf("Max: %.2f\n", maxTemp)
	fmt.Printf("Average: %.2f\n", averageTemp)
}
