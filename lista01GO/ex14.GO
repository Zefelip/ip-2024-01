package main

import (
    "fmt"
    "math"
)

func calcularAreaBase(a float64) float64 {
    return 3 * math.Pow(a, 2) * math.Sqrt(3) / 2
}

func calcularVolumePiramide(h, a float64) float64 {
    Ab := calcularAreaBase(a)
    return (1.0 / 3.0) * Ab * h
}

func main() {
    var h, a float64
    fmt.Scan(&h, &a)

    volume := calcularVolumePiramide(h, a)

    fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS\n", volume)
}