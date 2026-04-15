// Elabore um programa que apresente os resultados da soma e da média
// aritmética dos
// valores pares situados na
// faixa numérica de 50 a 70.

package main

import f "fmt"

func main() {

	var inicio int = 50
	var fim int = 70

	f.Print("Digite o início da faixa: ")
	f.Scan(&inicio)

	f.Print("Digite o fim da faixa: ")
	f.Scan(&fim)

	var soma int = 0
	var quantidade int = 0
	var media float64

	for i := inicio; i <= fim; i++ {
		if i%2 == 0 {
			soma += i
			quantidade++
		}
	}

	// Calcula a média
	media = float64(soma) / float64(quantidade)

	// Resultados
	f.Printf("Faixa numérica: %d a %d\n", inicio, fim)
	f.Printf("Soma dos valores pares: %d\n", soma)
	f.Printf("Quantidade de números pares: %d\n", quantidade)
	f.Printf("Média aritmética dos pares: %.2f\n", media)

}
