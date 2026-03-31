// Quadrado de pares - Att15
// Escreva um programa para ler um valor inteiro N e que gere o quadrado de cada um dos valores pares,
// de 1 até N, inclusive N, se for o caso.
// Entrada
// A entrada conterá uma linha com um valor inteiro N, 5 < N < 2000.
// Saída
// A saída deve conter, uma linha para cada quadrado computado. Em cada linha deve constar uma ex-
// pressão do tipo xˆ2 = y, onde x é um número par e y é o seu valor elevado ao quadrado. Imediatamente
// após o valor de y deve aparecer o caractere de quebra de linha: ‘\n’.

package main

import (
	"fmt"
)

func main() {
	var N int

	fmt.Print("Informe um número: ")
	fmt.Scan(&N)

	for i := 2; i <= N; i += 2 {
		quadrado := i * i
		fmt.Printf("%d^2 = %d\n", i, quadrado)
	}
}
