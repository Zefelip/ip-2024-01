/******************************************************************************
Faça um programa que leia vários pares de pontos, calcule o vetor definido entre eles e imprima a
coordenada do vetor que possui o maior valor absoluto (módulo). Considere que o vetor que liga dois
pontos A (x1,y1,z1) e B (x2,y2,z2) é calculada como: BA = (x2 − x1, y2 − y1, z2 − z1)
Entrada
A entrada consiste de várias linhas. A primeira linha apresenta um número de pontos N, com 2 ≤ N ≤
1000. As N linhas seguintes apresentam pontos no espaço na forma x y z, com x, y e z números reais tais
que -1000 ≤ x, y, z ≤ 1000. Faça um programa que calcule o vetor que liga dois pontos consecutivos nesta
lista e imprima a coordenada de maior valor absoluto. Note que, com exceção do primeiro e último valor
de entrada, todos os pontos serão utilizados duas vezes, uma para o cálculo do vetor com o ponto que veio
antes na lista e outra para o ponto que veio depois.
Saída
A saída consiste de (N-1) linhas, cada uma contendo o módulo do valor da coordenada de maior valor
absoluto, com 2 casas decimais após a vírgula. Após a impressão do último valor, quebre uma linha
*******************************************************************************/
package main

import (
	"fmt"
	"math"
)

func distancia(x1, x2, x3, y1, y2, y3 float64) float64 {
	distA := y1 - x1
	distB := y2 - x2
	distC := y3 - x3

	modulo := math.Pow(distA, 2) + math.Pow(distB, 2) + math.Pow(distC, 2)
	resultado := math.Sqrt(modulo)

	return resultado
}

func main() {
	var n int
	fmt.Println("Informe a quantidade de pontos a serem analisados:")
	fmt.Scanln(&n)

	a := make([]float64, n)
	b := make([]float64, n)
	c := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Informe o %dº ponto (x y z): ", i+1)
		fmt.Scanln(&a[i], &b[i], &c[i])
	}

	for i := 1; i < n; i++ {
		dist := distancia(a[i-1], b[i-1], c[i-1], a[i], b[i], c[i])
		fmt.Printf("Distância entre o ponto %d e o ponto %d: %.2f\n", i, i+1, dist)
	}
}
