mport (
    "fmt"
    "math"
)

func main() {
    var n int
    fmt.Scan(&n)

    if n <= 1 {
        fmt.Println("Numero invalido!")
        return
    }

    var soma float64
    for k := 1; k <= n; k++ {
        soma += 1.0 / float64(k)
    }

    fmt.Printf("%.6f\n", soma)
}