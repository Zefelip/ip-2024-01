package main

import "fmt"

func main() {
    var quantidade int
    fmt.Println("Digite a quantidade de times:")
    fmt.Scan(&quantidade)
    
    if quantidade < 2 {
        fmt.Println("Campeonato inválido!")
        return
    }
    
    final := 1
    
    for i := 1; i <= quantidade; i++ {
        for j := i + 1; j <= quantidade; j++ {
            fmt.Printf("Final %d: Time%d X Time%d\n", final, i, j)
            final++
        }
    }
}
