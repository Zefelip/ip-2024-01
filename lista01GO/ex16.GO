package main

import "fmt"

func main() {
    var salario float64
    fmt.Scan(&salario)

    var salarioReajustado float64

    if salario <= 300.00 {
        salarioReajustado = salario * 1.5 
    } else {
        salarioReajustado = salario * 1.3 
    }

    fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", salarioReajustado)
}
