/******************************************************************************
Escreva um programa em C para ler n elementos inteiros em um vetor, e depois imprimi-los.
Entrada
A entrada contém duas linhas. A primeira, contém um valor inteiro n < 5000 que corresponde ao
número de elementos que aparecem na segunda linha. A segunda linha contém n valores inteiros, separados
entre si por um espaço.
Saída
A saída é formada por uma linha contendo os n valores lidos.

*******************************************************************************/
package main

import "fmt"

func main() {
	var n int
	fmt.Println("Informe o tamanho do vetor")
	fmt.Scanln(&n)
	vetor := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&vetor[i])
	}
	fmt.Println("VETOR:", vetor)
}
