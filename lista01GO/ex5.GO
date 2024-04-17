package main

import "fmt"

func main() {
	var conta int
	var consumo float64
	var tipoConsumidor byte

	fmt.Scan(&conta, &consumo, &tipoConsumidor)

	var valorConta float64

	switch tipoConsumidor {
	case 'R':
		valorConta = 5.00 + consumo*0.05
	case 'C':
		if consumo <= 80 {
			valorConta = 500.00
		} else {
			valorConta = 500.00 + (consumo-80)*0.25
		}
	case 'I':
		if consumo <= 100 {
			valorConta = 800.00
		} else {
			valorConta = 800.00 + (consumo-100)*0.04
		}
	default:
		fmt.Println("Tipo de consumidor inválido.")
		return
	}

	fmt.Printf("CONTA = %d\n", conta)
	fmt.Printf("VALOR DA CONTA = %.2f\n", valorConta)
}
