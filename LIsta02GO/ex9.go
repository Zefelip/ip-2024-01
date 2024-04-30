package main
import "fmt"

func main() {
    var numero, auxiliar int 
    
    fmt.Println("Informe um número:")
    
    fmt.Scan(&numero)
    
    for i:=1; i<=numero; i++{
        if numero%i==0{
            auxiliar++
        }
    }
    
    if auxiliar == 2 {
        fmt.Println("É PRIMO")
    } else {
        fmt.Println("NÃO É PRIMO")
    }
}
