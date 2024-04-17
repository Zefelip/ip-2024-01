package main

import "fmt"

func fahrenheitParaCelsius(fahrenheit float64) float64 {
	return (5 * (fahrenheit - 32)) / 9
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var fahrenheit float64
		fmt.Scan(&fahrenheit)

		celsius := fahrenheitParaCelsius(fahrenheit)

		fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS\n", fahrenheit, celsius)
	}
}