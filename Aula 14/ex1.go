package main

import "fmt"

func potencia(x int, n int) int {
	if n == 0 {
		return 1
	}
	return x * potencia(x, n-1)
}

func main() {
	var n, x int
	fmt.Print("Informe o valor de x e de n:")
	fmt.Scan(&x, &n)

	fmt.Println(potencia(x, n))
}
