package main
import "fmt"

func main() {
    var m, n int
    fmt.Println("Digite a quantidade de linhas da matriz:")
    fmt.Scan(&m)
    fmt.Println("Digite a quantidade de colunas da matriz:")
    fmt.Scan(&n)
    
    for i:=0; i<m; i++{
        for j:=0; j<n; j++{
            if j<i {
                fmt.Printf("(%d-%d)", i+1, j+1)
            }
        }
        fmt.Println()
    }
}
