/******************************************************************************
Escreva um programa em C para armazenar n valores inteiros em um vetor, e depois imprimi-los na
ordem inversa a qual foram lidos.
Entrada
A entrada contém duas linhas. A primeira, contém um valor inteiro n < 5000 que corresponde ao
número de elementos que aparecem na segunda linha. A segunda linha contém n valores inteiros, separados
entre si por um espaço.
Saída
A saída é formada por uma linha contendo os n na ordem inversa da qual foram lidos.

*******************************************************************************/
package main

import "fmt"

func main() {
	var n int
	fmt.Println("Informe o tamanho do vetor:")
	fmt.Scanln(&n)
	for {
		if n > 5000 {
			fmt.Println("Valor Invalido")
			fmt.Scanln(&n)
		} else {
			break
		}
	}

	vetor := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Informe o elemento %d do vetor: ", i+1)
		fmt.Scanln(&vetor[i])
	}

	// Invertendo a ordem do vetor
	for i, j := 0, len(vetor)-1; i < j; i, j = i+1, j-1 {
		vetor[i], vetor[j] = vetor[j], vetor[i]
	}

	// Imprimindo o vetor invertido
	fmt.Println("Vetor invertido:", vetor)
}
