// Elabore um programa que leia dois números inteiros e efetue a adição dos mesmos. Caso o valor
// somado seja maior do que 20, o resultado a ser apresentado será a soma mencionada adicionada de
// 8. Caso o valor somado seja menor ou igual a 20, deve-se subtrair 5 do mesmo para apresentá-lo em
// seguida.

package main

import (
	"fmt"
)

func main() {

	var n1 int
	var n2 int
	var soma int

	fmt.Printf("Informe o primeiro número:")
	fmt.Scan(&n1)

	fmt.Printf("Informe o segundo número:")
	fmt.Scan(&n2)

	soma = n1 + n2

	if soma > 20 {
		soma = soma + 8
	} else {
		soma = soma - 5
	}

	fmt.Println("Resultado:", soma)

}
