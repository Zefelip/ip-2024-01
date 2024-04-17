package main
import "fmt"

func fahrenheitParaCelsius(fahrenheit float64) float64 {
	return (5 * (fahrenheit - 32)) / 9
}

func polegadaParaMilimetros(polegada float64) float64 {
    return (polegada*25.4)
}

func main() {
    var fahrenheit, polegada float64
    
    fmt.Scan(&fahrenheit, &polegada)
    
    celsius := fahrenheitParaCelsius(fahrenheit)
    milimetro := polegadaParaMilimetros(polegada)
    
    fmt.Printf("O VALOR EM CELSIUS = %.2f\nA QUANTIDADE DE CHUVA E = %.2f", celsius, milimetro)
}
