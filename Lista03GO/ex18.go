/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main

import "fmt"

func main() {
	var n int
	fmt.Println("Informe a quantidade de CPF's a serem analisados:")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var cpf [11]int
		for j := 0; j < len(cpf); j++ {
			fmt.Scan(&cpf[j])
		}

		soma1 := 0
		soma2 := 0

		// Calcula o primeiro dígito verificador (b1)
		for i := 0; i < 9; i++ {
			soma1 += cpf[i] * (i + 1)
		}
		b1 := soma1 % 11
		if b1 == 10 {
			b1 = 0
		}

		// Calcula o segundo dígito verificador (b2)
		for i := 0; i < 9; i++ {
			soma2 += cpf[i] * (9 - i)
		}

		b2 := soma2 % 11
		if b2 == 10 {
			b2 = 0
		}

		// Verifica se os dígitos verificadores estão corretos
		if b1 == cpf[9] && b2 == cpf[10] {
			fmt.Println("CPF valido")
		} else {
			fmt.Println("CPF invalido")
		}
	}
}
