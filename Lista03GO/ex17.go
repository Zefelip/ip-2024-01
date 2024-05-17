/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main

import "fmt"

func bubbleSort(vetor []int) {
	n := len(vetor)
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if vetor[j-1] > vetor[j] {
				vetor[j-1], vetor[j] = vetor[j], vetor[j-1]
			}
		}
	}
}

func main() {
	var n int
	var auxiliar int
	fmt.Scan(&n)

	vetor := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&vetor[i])
	}

	bubbleSort(vetor)

	for i := 2; i <= len(vetor); i++ {
		if vetor[i-2] != vetor[i-1] && vetor[i] != vetor[i-1] {
			auxiliar++
		}
	}
	fmt.Println(auxiliar)
}
