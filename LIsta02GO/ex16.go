package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)

	for h := 1; h <= n; h++ {
		for c1 := 1; c1 <= h; c1++ {
			for c2 := c1; c2 <= h; c2++ {
				if isPythagoreanTriple(c1, c2, h) {
					fmt.Printf("hipotenusa = %d, catetos %d e %d\n", h, c1, c2)
				}
			}
		}
	}
}

// Função para verificar se os números formam um triângulo retângulo
func isPythagoreanTriple(c1, c2, h int) bool {
	return c1*c1+c2*c2 == h*h
}