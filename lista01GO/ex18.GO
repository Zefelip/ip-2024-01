package main

import "fmt"

func main() {
    var a1, r, n int
    fmt.Scan(&a1, &r, &n)

    soma := 0
    for i := 0; i < n; i++ {
        termo := a1 + i*r
        soma += termo
    }

    fmt.Println(soma)
}
