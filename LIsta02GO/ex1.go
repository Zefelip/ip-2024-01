package main

import "fmt"

func notas(somaP float64, somaL float64, tf float64) float64 {
    return ((somaP / 8 * 0.7) + (somaL / 5 * 0.15) + tf*0.15) / 3
}

func main() {
    var p [8]float64
    var l [5]float64
    var tf float64
    var somaP, somaL float64
    var matricula int
    var horaAula int

    fmt.Scan(&matricula)

    for i := 0; i < 8; i++ {
        fmt.Scan(&p[i])
        somaP += p[i]
    }

    for i := 0; i < 5; i++ {
        fmt.Scan(&l[i])
        somaL += l[i]
    }

    fmt.Scan(&tf)

    media := notas(somaP, somaL, tf)

    fmt.Scan(&horaAula)

    var conceito string

    switch {
    case media >= 6.0 && horaAula >= 90:
        conceito = "Aprovado"
    case media < 6.0 && horaAula >= 90:
        conceito = "Reprovado por média"
    case media >= 6.0 && horaAula < 90:
        conceito = "Reprovado por falta"
    default:
        conceito = "Reprovado por média e falta"
    }

    fmt.Printf("Matrícula: %d, Nota Final: %.2f, Situação Final: %s", matricula, media, conceito)
}
