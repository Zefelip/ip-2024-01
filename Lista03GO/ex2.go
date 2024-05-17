/******************************************************************************
Dado um vetor V de tamanho N e um inteiro K, contabilize quantos elementos de V são maiores ou
iguais ao inteiro K.
Entrada
O programa terá apenas um caso de teste. O programa deve ler, obrigatoriamente, um número N que
pertença ao intervalo 1 ≤ N ≤ 1000. Se N lido não for válido, o programa deve fazer uma nova leitura de N.
Caso N seja válido, N representa o tamanho do vetor V. Na próxima linha há N números inteiros separados
por um espaço em branco cada, representando cada elemento do vetor V. E finalmente, na última linha há
um inteiro K.
Saída
Seu programa gera apenas uma linha de saída contendo um número inteiro representando quantos elementos do vetor V são maiores ou iguais ao inteiro K. Após a impressão do valor quebre uma linha.


*******************************************************************************/
package main

import "fmt"

func main() {
	var n, k, aux int
	fmt.Println("Informe o tamanho do vetor:")
	fmt.Scanln(&n)

	// Loop de validação para n
	for {
		if n < 1 || n > 1000 {
			fmt.Println("VALOR INVALIDO")
			fmt.Scanln(&n)
		} else {
			break
		}
	}

	vetor := make([]int, n)

	// Insere valores na fatia
	for i := 0; i < n; i++ {
		fmt.Scanln(&vetor[i])
	}

	fmt.Println("Informe o valor de k:")
	fmt.Scanln(&k)

	// Contagem de elementos maiores ou iguais a k
	for _, num := range vetor {
		if num >= k {
			aux++
		}
	}

	fmt.Printf("Número de elementos maiores ou iguais a %d: %d\n", k, aux)
}
