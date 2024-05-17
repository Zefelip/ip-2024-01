/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main

import (
	"fmt"
)

func main() {
	for {
		var n int
		fmt.Scanln(&n)

		if n == 0 {
			break // Termina quando N = 0
		}

		vetor := make([]int, n)

		// Ler os elementos do vetor
		for i := 0; i < n; i++ {
			fmt.Scanln(&vetor[i])
		}

		// Encontrar o maior elemento do vetor
		maior := 0
		for _, num := range vetor {
			if num > maior {
				maior = num
			}
		}

		// Criar um mapa para armazenar a contagem de frequência
		freq := make(map[int]int)

		// Iterar sobre os elementos do vetor e incrementar a contagem de cada valor no mapa
		for _, num := range vetor {
			for i := 0; i <= maior; i++ {
				if num <= i {
					freq[i]++
				}
			}
		}

		// Imprimir a contagem de frequência para cada valor menor ou igual ao maior elemento
		for i := 0; i <= maior; i++ {
			fmt.Printf("(%d) %d\n", i, freq[i])
		}
	}
}
