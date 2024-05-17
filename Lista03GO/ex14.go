/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main

import "fmt"

func bubbleSort(list []int) {
	n := len(list) // Corrigindo a função bubbleSort
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if list[j-1] > list[j] {
				list[j-1], list[j] = list[j], list[j-1]
			}
		}
	}
}

func main() {
	var n int
	var mediana float64
	fmt.Println("INFORME A QUANTIDADE DE NÚMEROS DE ENTRADA:")

	fmt.Scan(&n)

	vetor := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&vetor[i])
	}

	bubbleSort(vetor)

	if n%2 == 0 {
		mediana = float64(vetor[n/2]+vetor[(n/2)-1]) / 2 // Convertendo para float64
	} else {
		mediana = float64(vetor[n/2])
	}

	fmt.Println("Mediana:", mediana)
}
