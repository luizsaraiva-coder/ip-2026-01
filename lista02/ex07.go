// Escreva um programa que leia três valores inteiros distintos (assuma que o usuário digitará valores
// diferentes entre si) e os armazene nas variáveis A, B e C.
// Em seguida, descubra o menor valor e o armazene em uma variável denominada MENOR; o maior valor, coloque-o na variável MAIOR e o
// valor intermediário, na variável INTER.
// Imprima os valores em ordem crescente, ou seja, imprima as variáveis MENOR, INTER e MAIOR, nessa ordem.

package main

import (
	"fmt"
)

func main() {

	var A, B, C int
	var INTER, MENOR, MAIOR int

	fmt.Println("Determine os valores de A, B e C")
	fmt.Scan(&A, &B, &C)

	// O MENOR
	if A < B && A < C {
		MENOR = A
	} else if B < A && B < C {
		MENOR = B
	} else {
		MENOR = C
	}

	// O MAIOR
	if A > B && A > C {
		MAIOR = A
	} else if B > A && B > C {
		MAIOR = B
	} else {
		MAIOR = C
	}

	// O INTER
	if A != MENOR && A != MAIOR {
		INTER = A
	} else if B != MENOR && B != MAIOR {
		INTER = B
	} else {
		INTER = C
	}

	fmt.Println("Ordem crescente:")
	fmt.Println(MENOR, INTER, MAIOR)
}
