package main

import "fmt"

func main() {
    var fatorial, valor, quantidade int
    
    fmt.Scan(&fatorial)
    quantidade = fatorial
    
    valor = 1
    for fatorial > 1 {
        valor *= fatorial
        fatorial--
    }
    
    fmt.Printf("%d! = %d", quantidade, valor)
}
