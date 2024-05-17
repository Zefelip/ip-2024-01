/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main
import "fmt"

func bubbleSort (lista[] int){
    n := len(lista)
    for i:=0; i<n; i++{
        for j:=n-1; j>i; j--{
            if lista[j-1]>lista[j]{
                lista[j-1], lista[j] = lista[j], lista[j-1]
            }
        }
    }
}

func main() {
    var n int
    fmt.Println("Informe a quantidade de números que serão analisados:")
    fmt.Scan(&n)
    lista:=make([]int, n)
    
    for i:=0; i<n; i++{
        fmt.Scan(&lista[i])
    }
    
    bubbleSort(lista)
    
    fmt.Println("Números Ordenados:")
    for i:=0; i<len(lista);i++{
        fmt.Println(lista[i])
    }
}
