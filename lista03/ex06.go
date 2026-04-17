// Escreva um programa que receba um número inteiro positivo, verifique e informe se ele é ou não um número
// triangular. Obs.: Um número é triangular quando é resultado do produto de três números naturais consecutivos.

package main

import (
	"fmt"
	f "fmt"
)

func main() {
	var num int

	f.Print("Digite um número inteiro positivo: ")
	f.Scan(&num)

	if num <= 0 {
		fmt.Println("Número inválido.")
	} else {
		triangular := false

		for i := 1; i*(i+1)*(i+2) <= num; i++ {
			if i*(i+1)*(i+2) == num {
				triangular = true
				break

			}
		}

		if triangular {
			fmt.Println("É um número triangular.")
		} else {
			fmt.Println("Não é um número triangular.")
		}
	}
}
