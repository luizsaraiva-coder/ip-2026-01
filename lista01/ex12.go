// Locadora de charretes - Att12
// Uma locadora de charretes cobra R$ 10,00 de taxa para cada 3 horas de uso de uma charrete e R$5,00
// para cada 1 hora abaixo dessas 3 horas. Fazer um programa que leia a quantidade de horas que a charrete
// foi usada e que calcule e escreva quanto o cliente tem de pagar.
// Entrada
// O programa deve ler uma única linha na entrada, contendo o número de horas que o locatário utilizou a
// charrete.
// Saída
// O programa deve imprimir uma linha contendo a frase: O VALOR A PAGAR E = X, onde X é o valor
// do aluguel e deve conter no máximo 2 casas decimais. Após o valor do aluguel o programa deve imprimir
// um caractere de quebra de linha: ‘\n’.

package main

import (
	"fmt"
)

func main() {
	var horas int
	var valor float64

	fmt.Print("Informe a quantidade de horas: ")
	fmt.Scan(&horas)

	blocos := horas / 3
	resto := horas % 3

	valor = float64(blocos*10 + resto*5)

	fmt.Printf("O VALOR A PAGAR E = %.2f\n", valor)
}
