// Escrever um programa que leia o número de identificação, as 3 notas obtidas por um aluno nas 3 verificações e
// a média dos exercícios que fazem parte da avaliação. Calcular a média de aproveitamento do aluno, usando a
// fórmula:

// Média Final = ( nota1 + nota2 ∗ 2 + nota3 ∗ 3 + médiados exercícios )  / 7

// e o seu conceito, usando a tabela a seguir:
// Média de Aproveitamento                                              Conceito
// >= 9,0 e <= 10,0                                                       A
// >=7,5 e < 9,0                                                             B
// >=6,0 e < 7,5                                                           C
// >=4,0 e < 6,0                                                             D
// < 4,0                                                                   E

// O programa deve escrever o número do aluno, suas notas, a média dos exercícios, a média de aproveitamento,
// o conceito correspondente e a mensagem: APROVADO se o conceito for A, B ou C e REPROVADO, se o
// conceito for D ou E.

package main

import (
	"fmt"
)

func main() {
	var n1, n2, n3, mediaE, mediaF float64
	var conceito, aprovacao string
	var idAluno int

	fmt.Println("Informe o número de identificação do Aluno(a):")
	fmt.Scan(&idAluno)

	fmt.Println("Informe a nota 1, nota 2 e nota 3 do aluno(a):")
	fmt.Scan(&n1, &n2, &n3)

	fmt.Println("Informe a média dos exercícios do aluno")
	fmt.Scan(&mediaE)

	mediaF = (n1 + (n2 * 2) + (n3 * 3) + mediaE) / 7

	if mediaF >= 6 {
		aprovacao = "APROVADO"
	} else {
		aprovacao = "REPROVADO"
	}

	switch {
	case mediaF < 4:
		conceito = "E"
	case mediaF < 6:
		conceito = "D"
	case mediaF < 7.5:
		conceito = "C"
	case mediaF < 9:
		conceito = "B"
	case mediaF <= 10:
		conceito = "A"
	default:
		fmt.Println("Nota inválida!")
	}

	fmt.Printf("Identificação do Aluno(a): %d\n", idAluno)
	fmt.Printf("Notas:\nNota 1: %.2f Nota 2: %.2f Nota 3: %.2f\n", n1, n2, n3)
	fmt.Printf("Média Exercícios: %.2f\n", mediaE)
	fmt.Printf("Média de Aproveitamento: %.2f\n", mediaF)
	fmt.Println("Conceito: ", conceito)
	fmt.Println("Status: ", aprovacao)

}
