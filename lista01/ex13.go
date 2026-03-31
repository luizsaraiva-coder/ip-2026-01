// Conversão de Nota em Conceito - Att13
// Em um curso de mestrado as avaliações dos alunos no histórico escolar aparecem em forma de conceito.
// O regulamento do mestrado indica que um professor pode avaliar seus alunos com notas convencionais de
// zero a dez, mas precisa repassar à secretaria do curso a avaliação em termos de conceito. Nesse caso, a
// seguinte tabela de conversão deve ser usada pelo professor:
// Intervalo da Nota Conceito
// [9,0 a 10] A
// [7,5,9,0) B
// [6, 7,5) C
// [0 ,6,0) D

// Escreva um programa para ler um nota e converte-la no conceito correspondente.
// Entrada
// A entrada consiste de uma linha com um valor real entre 0 e 10 e com uma casa decimal.
// Saída
// O programa deve imprimir a seguinte frase: NOTA = x CONCEITO = y, onde x é o valor da nota lido
// na entrada, impresso com uma casa decimal y é o conceito correspondente. Após a frase, o programa deve
// imprimir o caractere de quebra de linha: ’\n’.

package main

import (
	"fmt"
)

func main() {
	var nota float64
	var conceito string

	fmt.Print("Informe a nota do aluno: ")
	fmt.Scan(&nota)

	if nota >= 9.0 && nota <= 10.0 {
		conceito = "A"
	} else if nota >= 7.5 && nota < 9.0 {
		conceito = "B"
	} else if nota >= 6.0 && nota < 7.5 {
		conceito = "C"
	} else if nota >= 0 && nota < 6.0 {
		conceito = "D"
	}

	fmt.Printf("NOTA = %.1f CONCEITO = %s\n", nota, conceito)
}
