package main
import "fmt"

func main() {
	var casos int
	fmt.Scan(&casos)

	for i := 1; i <= casos; i++ {
		var publicoTotal int
		var percentualPopular, percentualGeral, percentualArquibancada, percentualCadeiras float64
		fmt.Scan(&publicoTotal, &percentualPopular, &percentualGeral, &percentualArquibancada, &percentualCadeiras)

		rendaTotal := (float64(publicoTotal) * percentualPopular / 100 * 1.00) +
			(float64(publicoTotal) * percentualGeral / 100 * 5.00) +
			(float64(publicoTotal) * percentualArquibancada / 100 * 10.00) +
			(float64(publicoTotal) * percentualCadeiras / 100 * 20.00)

		fmt.Printf("A RENDA DO JOGO N. %d E = %.2f\n", i, rendaTotal)
	}
}
