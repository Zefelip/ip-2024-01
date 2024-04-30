package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    var numero [100]int // Define um array com tamanho máximo de 100
    for i := 0; i < n; i++ {
        fmt.Scan(&numero[i])
    }

    contador := 1 // Inicializa o contador com 1, pois o comprimento mínimo de um segmento crescente é 1
    aux := 1      // Inicializa aux com 1, pois o primeiro número já é um segmento crescente

    for i := 1; i < n; i++ {
        if numero[i-1] == numero[i] {
            aux++
            if aux > contador {
                contador = aux
            }
        } else {
            aux = 1 // Reseta aux para 1 quando o segmento crescente termina
        }
    }

    fmt.Printf("A maior subsequência de elementos iguais possui %d elementos", contador)
}
