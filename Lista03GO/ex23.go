package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Split(input, ";")
	if len(parts) != 2 {
		fmt.Println("FORMATO INVALIDO!")
		return
	}

	A := parts[0]
	B := parts[1]

	freqA := vowelFrequency(A)
	freqB := vowelFrequency(B)

	fmt.Printf("(%d,%d,%d,%d,%d)\n", freqA[0], freqA[1], freqA[2], freqA[3], freqA[4])
	fmt.Printf("(%d,%d,%d,%d,%d)\n", freqB[0], freqB[1], freqB[2], freqB[3], freqB[4])

	distance := euclideanDistance(freqA, freqB)
	fmt.Printf("%.2f\n", distance)
}

func vowelFrequency(s string) [5]int {
	s = strings.ToLower(s)
	var freq [5]int
	for _, char := range s {
		switch char {
		case 'a':
			freq[0]++
		case 'e':
			freq[1]++
		case 'i':
			freq[2]++
		case 'o':
			freq[3]++
		case 'u':
			freq[4]++
		}
	}
	return freq
}

func euclideanDistance(a, b [5]int) float64 {
	var sum float64
	for i := 0; i < 5; i++ {
		diff := a[i] - b[i]
		sum += float64(diff * diff)
	}
	return math.Sqrt(sum)
}
