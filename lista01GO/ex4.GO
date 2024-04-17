package main

import "fmt"

func main() {
    var kw, salarioMin float64
    var custoKw, consumo, custoDesconto float64
    
    fmt.Scan(&kw, &salarioMin)
    
    custoKw = 70.0 / 10000.0 * salarioMin
    
    consumo = custoKw * kw
    
    custoDesconto = consumo * 90.0 / 100.0
    
    fmt.Printf("Custo por kw: %.2f\nCusto do consumo: %.2f\nCusto com desconto: %.2f", custoKw, consumo, custoDesconto)
}
