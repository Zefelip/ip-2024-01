package main

import (
    "fmt"
    "math"
)

func areaTotalCilindro(areaCilindro float64, areaLateralCilindro float64) float64 {
    return 2*areaCilindro + areaLateralCilindro
}

func areaCilindro(raio float64) float64 {
    return math.Pi * raio * raio
}

func areaLateralCilindro(raio float64, altura float64) float64 {
    return 2 * math.Pi * raio * altura
}

func custo(areaTotalCilindro float64) float64 {
    return areaTotalCilindro * 100
}

func main() {
    var raio, altura float64

    fmt.Scan(&raio)
    fmt.Scan(&altura)

    areaLateral := areaLateralCilindro(raio, altura)
    area := areaCilindro(raio)
    areaTotal := areaTotalCilindro(area, areaLateral)
    custo := custo(areaTotal)

    fmt.Printf("O VALOR DO CUSTO É: %.2f", custo)
}
