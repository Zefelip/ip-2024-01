package main

import "fmt"

func main() {
	var T int
	fmt.Scanln(&T)

	for i := 0; i < T; i++ {
		var A, B int
		fmt.Scanln(&A, &B)

		// Inverte os números para compará-los corretamente
		A = reverseNumber(A)
		B = reverseNumber(B)

		if A > B {
			fmt.Println(A)
		} else {
			fmt.Println(B)
		}
	}
}

// Função para inverter o número
func reverseNumber(num int) int {
	var reversedNum int
	for num > 0 {
		digit := num % 10
		reversedNum = reversedNum*10 + digit
		num /= 10
	}
	return reversedNum
}
