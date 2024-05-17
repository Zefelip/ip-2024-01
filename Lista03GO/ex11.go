/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main

import "fmt"

func main() {
	var n int
	fmt.Println("Informe a quantidade de números a serem analisados:")
	fmt.Scan(&n)

	for {
		if n < 1 || n > 1000 {
			fmt.Println("Valor Invalido")
			fmt.Scan(&n)
		} else {
			break
		}
	}
	vetor := make([]int, n)
	vetorInverso := make([]int, n)
	menor := 0
	maior := 0

	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])

		if i == 0 || maior < vetor[i] {
			maior = vetor[i]
		}

		if i == 0 || menor > vetor[i] {
			menor = vetor[i]
		}
	}

	for i := 0; i < n; i++ {
		vetorInverso[i] = vetor[n-i-1] // Corrigido: preenchendo o vetor inverso corretamente
	}

	fmt.Println(vetor)
	fmt.Println(vetorInverso)
	fmt.Println(maior)
	fmt.Println(menor)
}
