package main

import "fmt"

func main() {
    var populacaoA, populacaoB float64
    var anos int
    fmt.Scan(&populacaoA, &populacaoB)

    for populacaoB > populacaoA {
        populacaoA = populacaoA + populacaoA*0.03
        populacaoB = populacaoB + populacaoB*0.015
        anos++
    }

    fmt.Printf("ANOS = %d", anos)
}