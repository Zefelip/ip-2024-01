/******************************************************************************
A Mega-Sena é a maior loteria do Brasil. Para ganhar o prêmio máximo é necessário acertar a sena,
o que significa obter coincidência entre seis dos números apostados e os seis números sorteados, de um
total de sessenta dezenas (de 01 a 60), independentemente da ordem da aposta ou da ordem do sorteio.
O concurso prevê também a chance de ganhar parte do prêmio, acertando a quina ou a quadra. A MegaSena foi lançada em março de 1996 e já premiou mais de 200 ganhadores na faixa principal. Os prêmios
correspondem a 32,2% da renda das apostas ao imposto de renda correspondem 13,8% de todas as apostas.
Os vencedores têm 90 dias para retirar o prêmio, se o período expirar, o dinheiro do prêmio será transferido
ao Tesouro Nacional e investido em programas educacionais. Vale lembrar que a probabilidade de acerto
em uma única aposta de 6 dezenas é de 1 em 50.063.860, o que representa um percentual de 0,000002%.
Faça um programa que receba todas as apostas e as seis dezenas sorteadas de um concurso e mostre quantos
vencedores para sena, quina e quadra houve.
Entrada
Na primeira linha da entrada haverá uma linha com as seis dezenas sorteadas, separadas por um espaço
em branco cada. Na linha seguinte haverá um inteiro N, 1 ≤ N ≤ 50000, representando a quantidade de
apostas. Em seguida, em cada uma das N linhas haverá as seis dezenas de cada aposta, sendo que as dezenas
estão no intervalo entre 1 e 60 e sem repetição de dezenas por apostas.
Saída
A saída consiste de 3 linhas contando uma das seguintes frases: “Houve K acertador(es) da sena” ou
“Houve K acertador(es) da quina” ou ainda “Houve K acertador(es) da quadra”, onde K é quantidade de
acertadores para a faixa. Caso não haja acertadores a seguinte frase deve ser apresentada: “Nao houve
acertador para sena” ou “Nao houve acertador para quina” ou ainda “Nao houve acertador para quadra”.
Ao exibir a última frase quebre uma linha.
*******************************************************************************/
package main

import (
	"fmt"
	"sort"
)

func main() {
	var numeros [6]int
	var jogos [][6]int
	var n int
	var sena, quina, quadra int

	fmt.Println("Informe os números sorteados na Mega-Sena:")
	for i := 0; i < len(numeros); i++ {
		fmt.Scan(&numeros[i])
	}

	sort.Ints(numeros[:]) // Convertendo array para slice e ordenando

	fmt.Println("Informe a quantidade de jogos a serem analisados:")
	fmt.Scanln(&n)

	jogos = make([][6]int, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Informe os números do %dº jogo:\n", i+1)
		for j := 0; j < 6; j++ {
			fmt.Scan(&jogos[i][j])
		}
		sort.Ints(jogos[i][:]) // Convertendo array para slice e ordenando
	}

	for i := 0; i < n; i++ {
		acertos := 0
		for _, num := range jogos[i] {
			if contains(numeros[:], num) {
				acertos++
			}
		}
		switch acertos {
		case 6:
			sena++
		case 5:
			quina++
		case 4:
			quadra++
		}
	}

	fmt.Printf("HOUVE %d ACERTADORES DA SENA\n", sena)
	fmt.Printf("HOUVE %d ACERTADORES DA QUINA\n", quina)
	fmt.Printf("HOUVE %d ACERTADORES DA QUADRA\n", quadra)
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
