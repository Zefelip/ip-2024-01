package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Função que inverte uma palavra
func reverseWord(word string) string {
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Lê uma linha da entrada
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()

		// Condição de parada: linha vazia
		if strings.TrimSpace(line) == "" {
			break
		}

		// Divide a linha em palavras
		words := strings.Fields(line)

		// Inverte cada palavra
		for i, word := range words {
			words[i] = reverseWord(word)
		}

		// Inverte a ordem das palavras
		for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
			words[i], words[j] = words[j], words[i]
		}

		// Concatena as palavras invertidas em uma frase
		reversedLine := strings.Join(words, " ")

		// Imprime a linha invertida
		fmt.Println(reversedLine)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Erro ao ler a entrada:", err)
	}
}
