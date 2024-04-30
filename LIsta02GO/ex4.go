package main

import "fmt"

func main() {
    var n, K, i, s float64
    var a, b, c, d float64
    var soma, subtracao, multiplicacao, divisao float64
    
    fmt.Scan(&n, &K, &i, &s)
    
    a=K
    b=K
    c=K
    d=K
    
    fmt.Println("Tabuada de Soma:")
    for j := 1; j <= int(i); j++ {
        soma = n + a
        fmt.Printf("%.2f + %.2f = %.2f\n", n, a, soma)
        a += s
    }
    
    fmt.Println("Tabuada de Subtração:")
    for j := 1; j <= int(i); j++ {
        subtracao = n - b
        fmt.Printf("%.2f - %.2f = %.2f\n", n, b, subtracao)
        b += s
    }
    
    fmt.Println("Tabuada de Multiplicação:")
    for j := 1; j <= int(i); j++ {
        multiplicacao = n * c
        fmt.Printf("%.2f + %.2f = %.2f\n", n, c, multiplicacao)
        c += s
    }
    
    fmt.Println("Tabuada de Divisão:")
    for j := 1; j <= int(i); j++ {
        divisao = n / d
        fmt.Printf("%.2f + %.2f = %.2f\n", n, d, divisao)
        d += s
    }
}
