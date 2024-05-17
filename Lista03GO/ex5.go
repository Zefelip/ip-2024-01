/******************************************************************************
Faça um programa que receba vários vetores e informe para cada um deles qual o maior elemento e o
indice (da primeira ocorrência) em que encontra-se tal elemento.
Entrada
O programa possui vários casos de testes. A primeira de cada caso contem um inteiro N, 1 < N ≤ 10000,
representando o tamanho do vetor. A segunda linha conterá N inteiros entre 0 e 1000, representando os N
elementos do vetor. A entrada termina quando N=0.
Saída
O programa gera N linhas de saída, com dois inteiros separados por um espaço em branco. O primeiro
inteiro é o índice da primeira ocorrência do maior elemento do vetor e o segundo inteiro é o maior valor do
vetor. Após a impressão de cada saída, inclusive a última, quebre uma linha.

*******************************************************************************/
package main

import "fmt"

func main() {
	var n int
	fmt.Println("Informe a quantidade de vetores a serem analisados")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var m, valor, indice int
		fmt.Printf("Informe a quantidade de caracteres do vetor %d: ", i+1)
		fmt.Scanln(&m)

		for {
			if m < 0 || m > 1000 {
				fmt.Println("VALOR INVALIDO")
				fmt.Scanln(&m)
			} else {
				break
			}
		}

		vetor := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scanln(&vetor[j])
		}

		valor = vetor[0]
		for j, num := range vetor {
			if num > valor {
				valor = num
				indice = j
			}
		}

		fmt.Printf("Indice: %d\nValor:%d\n", indice, valor)
	}
}
