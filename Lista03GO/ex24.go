package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countingSort(arr []int, maxVal int) []int {
	count := make([]int, maxVal+1)
	sorted := make([]int, len(arr))

	// Passo 3: Contagem de cada elemento
	for _, num := range arr {
		count[num]++
	}

	// Passo 4: Acumulação
	for i := 1; i <= maxVal; i++ {
		count[i] += count[i-1]
	}

	// Passo 5: Construção do array ordenado
	for i := len(arr) - 1; i >= 0; i-- {
		num := arr[i]
		sorted[count[num]-1] = num
		count[num]--
	}

	return sorted
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Leitura de N
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		N, err := strconv.Atoi(line)
		if err != nil || N == 0 {
			break
		}

		// Leitura do array
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		elements := strings.Split(line, " ")
		arr := make([]int, N)

		maxVal := 0
		for i, el := range elements {
			num, _ := strconv.Atoi(el)
			arr[i] = num
			if num > maxVal {
				maxVal = num
			}
		}

		// Ordenação com Counting Sort
		sortedArr := countingSort(arr, maxVal)

		// Impressão do resultado
		for i, num := range sortedArr {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(num)
		}
		fmt.Println()
	}
}
