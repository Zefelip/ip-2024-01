/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Println("Informe a quantidade de valores a serem lidos:")
	fmt.Scan(&n)

	for {
		if n < 2 || n > 1000 {
			fmt.Println("VALOR INVÁLIDO")
			fmt.Scan(&n)
		} else {
			break
		}
	}

	for i := 0; i <= n/2; i++ {
		var a, b, c float64
		var x, y, z float64
		fmt.Scan(&a, &b, &c)
		fmt.Scan(&x, &y, &z)

		aux := (math.Pow(a-x, 2)) + (math.Pow(b-y, 2)) + (math.Pow(c-z, 2))
		distancia := math.Sqrt(aux)
		fmt.Printf("A distância entre os pontos é: %.2f", distancia)
	}
}
