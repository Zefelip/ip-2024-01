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
	fmt.Println("Informe a quantidade de números a serem lidos:")

	fmt.Scan(&n)

	vetor := make([]int, n)
	maiorNota := 0
	posicaoPrimeiraOcorrencia := -1
	ultimaNota := 0
	frequenciaUltimaNota := 0

	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])

		if vetor[i] > maiorNota {
			maiorNota = vetor[i]
			posicaoPrimeiraOcorrencia = i
		}

		if i == n-1 {
			ultimaNota = vetor[i]
		}
	}

	for _, num := range vetor {
		if num == ultimaNota {
			frequenciaUltimaNota++
		}
	}

	fmt.Printf("Nota %d, %d vezes\n", ultimaNota, frequenciaUltimaNota)
	fmt.Printf("Nota %d, indice %d", maiorNota, posicaoPrimeiraOcorrencia)
}
