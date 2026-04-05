// Faça um programa que indique se um número inteiro informado pelo usuário está compreendido
// entre 20 e 90 ou não. (20 e 90 não estão na faixa de valores).

package main

import (
	"fmt"
)

func main() {

	var n int

	fmt.Println("Digite um número inteiro")
	fmt.Scan(&n)

	if n > 20 && n < 90 {

		fmt.Println("O número está entre 20 e 90")

	} else {

		fmt.Println("O número não está entre 20 e 90")
	}
}
