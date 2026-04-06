// Uma locadora tem as seguintes regras para aluguel de DVDs:
// - Às segundas, terças e quintas (2, 3 e 5) : desconto de 40% em relação ao preço normal;
// - Às quartas , sextas, sábados e domingos (4,6 ,7 e 1): preço normal;
// - Aluguel de DVDs comuns: preço normal;
// - Aluguel de lançamentos: acréscimo de 15% em relação ao preço normal.
// Desenvolver um programa para ler o preço normal do DVD alugado (em R$) e sua categoria (comum ou
// lançamento). Calcular e imprimir o preço final que será pago pela locação do DVD.

package main

import (
	"fmt"
	"strings"
)

func main() {
	categoriaDvd := ""
	diaSemana := 0
	precoB, precoF := 0.0, 0.0

	fmt.Println("--- Sistema de Locadora ---")
	fmt.Print("Preço normal do DVD (R$): ")
	fmt.Scan(&precoB)
	fmt.Print("Categoria (C - Comum / L - Lançamento): ")
	fmt.Scan(&categoriaDvd)
	fmt.Print("Dia da semana (1-Dom, 2-Seg, 3-Ter, 4-Qua, 5-Qui, 6-Sex, 7-Sab): ")
	fmt.Scan(&diaSemana)

	categoriaDvd = strings.ToUpper(categoriaDvd)

	precoF = precoB

	if categoriaDvd == "L" {
		precoF *= 1.15
	} else if categoriaDvd != "L" {
		fmt.Println("Categoria inválida! Considerando como Comum.")
	}

	switch diaSemana {
	case 4, 6, 7, 1:
		precoF *= 0.60
	case 2, 3, 5:
	default:
		fmt.Println("Dia da semana inválido!")
		return
	}
	fmt.Printf("Valor total a ser pago é %.2f\n", precoF)

}
