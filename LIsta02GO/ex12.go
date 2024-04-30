package main

import "fmt"

func main() {
    var ValorIngresso, ValorInicial, ValorFinal float64

    fmt.Scan(&ValorIngresso, &ValorInicial, &ValorFinal)

    if ValorInicial >= ValorFinal {
        fmt.Println("INTERVALO INVALIDO")
        return
    }

    var melhorValorFinal, melhorLucro float64
    var melhorNumIngressos int

    for preco := ValorInicial; preco <= ValorFinal; preco++ {
        var numIngressos int
        var lucro float64

        if preco < ValorIngresso {
            numIngressos = 120 + 25*(int(ValorIngresso)-int(preco))
        } else {
            numIngressos = 120 - 30*(int(preco)-int(ValorIngresso))
        }

        if numIngressos <= 0 {
            lucro = -200
        } else {
            lucro = float64(numIngressos)*preco - (0.05*float64(numIngressos) + 200)
        }

        fmt.Printf("V: %.2f, N: %d, L: %.2f\n", preco, numIngressos, lucro)

        if lucro > melhorLucro {
            melhorLucro = lucro
            melhorValorFinal = preco
            melhorNumIngressos = numIngressos
        }
    }

    fmt.Printf("Melhor valor final: %.2f\nLucro: %.2f\nNumero de ingressos: %d\n", melhorValorFinal, melhorLucro, melhorNumIngressos)
}