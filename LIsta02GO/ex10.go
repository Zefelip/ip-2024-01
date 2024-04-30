package main
import "fmt"

func salario (horas int, valorHora float64) float64 {
    return float64(horas) * valorHora
}

func main() {
    var quantidade int
    
    fmt.Println("Qual a quantidade de linhas a serem analisadas?")
    fmt.Scan(&quantidade)
    
    matricula:= make([]int, quantidade)
    horas:=make([]int, quantidade)
    valorHora:=make([]float64, quantidade)
    
    for i:=0; i<quantidade; i++{
        fmt.Scan(&matricula[i], &horas[i], &valorHora[i])
    }
    
    for i:=0; i<quantidade; i++{
        salario:= salario(horas[i], valorHora[i])
        fmt.Printf("MATRICULA:%d\nSALARIO:%.2f", matricula[i], salario)
    }
    
    
}
