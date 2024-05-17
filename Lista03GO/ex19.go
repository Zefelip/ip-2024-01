/******************************************************************************
Dada uma sequência de N números inteiros na ordem crescente, identifique os elementos únicos.
Entrada
Na primeira linha há um inteiro N, 1 ≤ N ≤ 1000, representando a quantidade de elementos. Nas N
linhas seguintes haverá um número inteiro (positivo ou negativo) por linha.
Saída
O programa apresenta um sequência de elementos únicos, na ordem crescente. Cada elemento é apresentado em uma linha

*******************************************************************************/
package main

import "fmt"

func bubbleSort(vetor []int) {
	n := len(vetor)
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if vetor[j-1] > vetor[j] {
				vetor[j-1], vetor[j] = vetor[j], vetor[j-1]
			}
		}
	}
}

func main() {
	var n int
	fmt.Println("Informe a quantidade de números a serem analisados:")
	fmt.Scanln(&n)

	vetor := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&vetor[i])
	}

	bubbleSort(vetor)

	auxiliar := make([]int, len(vetor))

	for i := 1; i < n; i++ {
		if i == 1 && vetor[i-1] < vetor[i] {
			auxiliar[i] = vetor[i-1]
		} else if vetor[i] > vetor[i-1] {
			auxiliar[i] = vetor[i]
		} else {
			auxiliar[i] = 0
		}
	}

	for i := 0; i < n; i++ {
		if auxiliar[i] != 0 {
			fmt.Println(auxiliar[i])
		}
	}
}
