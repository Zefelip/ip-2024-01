package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Função que ordena um vetor de inteiros em ordem crescente
func ordenaVetor(vetor []int) {
	sort.Ints(vetor)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Solicita o tamanho do vetor
	fmt.Print("Digite o tamanho do vetor: ")
	scanner.Scan()
	tamanhoStr := scanner.Text()
	tamanho, err := strconv.Atoi(tamanhoStr)
	if err != nil {
		fmt.Println("Erro ao ler o tamanho do vetor:", err)
		return
	}

	// Solicita os elementos do vetor
	fmt.Printf("Digite %d números inteiros separados por espaço: ", tamanho)
	scanner.Scan()
	elementosStr := scanner.Text()

	// Converte os elementos para inteiros
	elementos := strings.Fields(elementosStr)
	if len(elementos) != tamanho {
		fmt.Println("Número de elementos não corresponde ao tamanho informado.")
		return
	}

	vetor := make([]int, tamanho)
	for i, elem := range elementos {
		vetor[i], err = strconv.Atoi(elem)
		if err != nil {
			fmt.Println("Erro ao converter elemento para inteiro:", err)
			return
		}
	}

	// Ordena o vetor
	ordenaVetor(vetor)

	// Exibe o vetor ordenado
	fmt.Println("Vetor ordenado:", vetor)
}
