package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Função que troca dois elementos de um vetor entre si
func troca(vetor []int, i, j int) {
	vetor[i], vetor[j] = vetor[j], vetor[i]
}

// Função que troca elementos opostos se a condição for atendida
func trocaOpostosSeMenor(vetor []int, tamanho int) {
	for i := 0; i < tamanho/2; i++ {
		oposto := tamanho - 1 - i
		if vetor[i] < vetor[oposto] {
			troca(vetor, i, oposto)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Lê o número de casos de teste
	scanner.Scan()
	numCasos, _ := strconv.Atoi(scanner.Text())

	// Processa cada caso de teste
	for caso := 0; caso < numCasos; caso++ {
		// Lê o tamanho do vetor
		scanner.Scan()
		tamanho, _ := strconv.Atoi(scanner.Text())

		// Lê os elementos do vetor
		scanner.Scan()
		elementos := strings.Fields(scanner.Text())

		// Converte os elementos para inteiros
		vetor := make([]int, tamanho)
		for i, elem := range elementos {
			vetor[i], _ = strconv.Atoi(elem)
		}

		// Processa o vetor
		trocaOpostosSeMenor(vetor, tamanho)

		// Imprime o vetor processado
		for i, val := range vetor {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	}
}
