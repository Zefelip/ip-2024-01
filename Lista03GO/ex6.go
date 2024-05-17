/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main

import "fmt"

func main() {
	var n, soma int
	fmt.Scanln(&n)
	vetor := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&vetor[i])
		soma += vetor[i]
	}

	fmt.Println(soma)
}
