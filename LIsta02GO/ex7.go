
package main

import "fmt"

func main() {
    var quantidade, impar, par, quantidadeImpar, quantidadePar int
    var mediaPar, mediaImpar float64
    
    fmt.Scan(&quantidade)
    
    // Handle division by zero
    if quantidade == 0 {
        fmt.Println("Cannot calculate average for zero numbers.")
        return
    }
    
    // Initialize variables
    mediaPar = 0
    mediaImpar = 0
    
    for i := 0; i < quantidade; i++ {
        var numero int
        fmt.Scan(&numero)
        if numero%2 == 0 {
            // Increment par by the value of numero
            par += numero
            quantidadePar++
        } else {
            // Increment impar by the value of numero
            impar += numero
            quantidadeImpar++
        }
    }
    
    // Calculate averages
    mediaPar = float64(par) / float64(quantidadePar)
    mediaImpar = float64(impar) / float64(quantidadeImpar)
    
    fmt.Printf("MEDIA PAR: %.2f\nMEDIA IMPAR: %.2f", mediaPar, mediaImpar)
}
