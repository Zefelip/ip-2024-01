package main

import (
	"fmt"
	"strings"
)

func main() {
	for {
		var n, d int
		fmt.Scan(&n, &d)

		if n == 0 && d == 0 {
			break
		}

		var number string
		fmt.Scan(&number)

		maxPrize := calculateMaxPrize(number, d)
		fmt.Println(maxPrize)
	}
}

func calculateMaxPrize(number string, d int) string {
	var stack []byte

	for i := range number {
		for len(stack) > 0 && stack[len(stack)-1] < number[i] && len(stack)+len(number)-i > len(number)-d {
			stack = stack[:len(stack)-1]
		}

		if len(stack) < len(number)-d {
			stack = append(stack, number[i])
		}
	}

	return strings.Join(strings.Split(string(stack), ""), "")
}
