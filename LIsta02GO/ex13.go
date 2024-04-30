package main

import "fmt"

func main() {
    var n, soma int 
    var quantidadeGraos [][]int // Removido o tamanho da matriz
    
    // Solicita o número de grãos para a posição [1][1]
    fmt.Println("Escreva o número de grãos para a posição [1][1]: ")
    fmt.Scan(&n)
    
    // Inicializa a matriz com 8 linhas e 8 colunas
    quantidadeGraos = make([][]int, 8)
    for i := range quantidadeGraos {
        quantidadeGraos[i] = make([]int, 8)
    }
    
    // Atribui o valor de n à posição [1][1]
    quantidadeGraos[0][0] = n
    
    // Preenche a matriz e calcula a soma dos valores
    for i := 0; i < 8; i++ {
        for j := 0; j < 8; j++ {
            if j%2 != 0 {
                quantidadeGraos[i][j] = n
            } else {
                quantidadeGraos[i][j] = 2 * n
            }
            soma += quantidadeGraos[i][j]
        }
    }
    
    // Imprime a quantidade total de grãos
    fmt.Printf("Quantidade de Grãos: %d\n", soma)
}
