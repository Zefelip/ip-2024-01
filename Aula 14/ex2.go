package main
import "fmt"

func main() {
    var n int
    fmt.Print("Informe o tamanho do array:")
    
    fmt.Scan(&n)
    arr := make([]float64, n)
    
    fmt.Print("Informe os valores:")
    for i:=0; i<len(arr); i++{
        fmt.Scan(&arr[i])
    }
    
    fmt.Print(soma(arr))
}

func soma(arr []float64) float64 {
    if len(arr) == 0 {
        return 0
    }
    return arr[0] + soma(arr[1:])
}