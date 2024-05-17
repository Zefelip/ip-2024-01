/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main

import "fmt"

func main() {
	var N int
	fmt.Scanln(&N)

	// Mapa para armazenar a frequência dos números
	frequencias := make(map[int]int)

	// Leitura dos números e contagem da frequência
	for i := 0; i < N; i++ {
		var num int
		fmt.Scanln(&num)
		frequencias[num]++
	}

	// Encontrar o valor com a maior frequência
	var maiorFrequencia, valorMaiorFrequencia int
	for num, freq := range frequencias {
		if freq > maiorFrequencia || (freq == maiorFrequencia && num < valorMaiorFrequencia) {
			maiorFrequencia = freq
			valorMaiorFrequencia = num
		}
	}

	// Saída
	fmt.Println(valorMaiorFrequencia)
	fmt.Println(maiorFrequencia)
}
