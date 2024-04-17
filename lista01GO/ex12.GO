package main

import "fmt"

func main() {
	var horas int
	fmt.Scan(&horas)

	var valorAPagar float64

	if horas >= 3 {
		valorAPagar = float64(horas/3) * 10.00
		if horas%3 != 0 {
			valorAPagar += 5.00
		}
	} else {
		valorAPagar = float64(horas) * 5.00
	}

	fmt.Printf("O VALOR A PAGAR E = %.2f\n", valorAPagar)
}
