package main

import "fmt"

// Função para calcular o MDC (Máximo Divisor Comum)
func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

// Função para calcular o MMC (Mínimo Múltiplo Comum)
func lcm(a, b int) int {
    return a * b / gcd(a, b)
}

// Função para calcular o MMC de três números
func lcm3(a, b, c int) int {
    return lcm(a, lcm(b, c))
}

func main() {
    var num1, num2, num3 int

    // Leitura dos três números inteiros diferentes de zero
    fmt.Println("Digite três números inteiros diferentes de zero:")
    fmt.Scan(&num1, &num2, &num3)

    // Cálculo do MMC dos três números
    result := lcm3(num1, num2, num3)

    // Saída conforme especificado
    fmt.Printf("%5d %5d %5d :%d\n", num1, num2, num3, result)
}
