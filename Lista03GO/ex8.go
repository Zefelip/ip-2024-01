/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func decimalToBinary(decimal int) string {
	// Converte o número decimal para binário
	binary := ""
	for decimal > 0 {
		remainder := decimal % 2
		binary = strconv.Itoa(remainder) + binary
		decimal /= 2
	}
	return binary
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// Lê a entrada como string
		input := scanner.Text()

		// Converte a string para inteiro
		decimal, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Erro ao converter entrada para inteiro:", err)
			return
		}

		// Converte o número decimal para binário
		binary := decimalToBinary(decimal)

		// Imprime o número binário
		fmt.Println(binary)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler entrada:", err)
	}
}
