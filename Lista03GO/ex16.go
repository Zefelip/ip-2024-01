/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main

import "fmt"

func inverterVetor(vetor []int) {
	tamanho := len(vetor)
	for i := 0; i < tamanho/2; i++ {
		// Trocar o elemento atual com o elemento simétrico
		vetor[i], vetor[tamanho-1-i] = vetor[tamanho-1-i], vetor[i]
	}
}

func main() {
	var qntdAlunos, presencaMin int
	var auxiliar int
	fmt.Println("Informe, respectivamente, a quantidade de alunos matriculados na turma e a quantidade mínima de presença exigida")
	fmt.Scan(&qntdAlunos, &presencaMin)

	alunos := make([]int, qntdAlunos)

	for i := 0; i < qntdAlunos; i++ {
		fmt.Scanln(&alunos[i])
	}

	for _, i := range alunos {
		if i <= 0 {
			auxiliar++
		}
	}

	if auxiliar < presencaMin {
		fmt.Println("SIM")
	} else {
		fmt.Println("NÃO")
		for i := qntdAlunos - 1; i >= 0; i-- {
			if alunos[i] <= 0 {
				fmt.Println(i + 1)
			}
		}
	}

}
