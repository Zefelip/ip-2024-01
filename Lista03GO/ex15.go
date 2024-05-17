package main

import "fmt"

// Função para calcular a soma dos termos de uma PA
func somaPA(n int) int {
	return (n - 1) * n / 2
}

// Função para ordenar um vetor usando o algoritmo Bubble Sort
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
	fmt.Println("Informe a quantidade de PAs a serem analisadas:")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var termos int
		fmt.Println("Informe a quantidade de termos da PA:")
		fmt.Scanln(&termos)
		vetor := make([]int, termos)

		soma := somaPA(termos)

		fmt.Println("Informe os termos da PA:")
		for j := 0; j < termos; j++ {
			fmt.Scanln(&vetor[j])
		}

		bubbleSort(vetor)

		menor := vetor[1] - vetor[0] // Inicialize menor com a diferença entre os dois primeiros termos

		for j := 2; j < termos; j++ {
			if vetor[j]-vetor[j-1] < menor {
				menor = vetor[j] - vetor[j-1]
			}
		}

		fmt.Printf("Menor distância: %d\nNúmero de comparações: %d\n", menor, soma)
	}
}
