package main

import "fmt"

func main() {
    var nota float64
    fmt.Scan(&nota)

    var conceito string

    switch {
    case nota >= 9.0 && nota <= 10:
        conceito = "A"
    case nota >= 7.5 && nota < 9.0:
        conceito = "B"
    case nota >= 6.0 && nota < 7.5:
        conceito = "C"
    default:
        conceito = "D"
    }

    fmt.Printf("NOTA = %.1f CONCEITO = %s\n", nota, conceito)
}