// Crie um programa que leia um valor inteiro e diga se ele é positivo, negativo ou nulo.

package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Informe um número:")
	fmt.Scan(&n)

	if n > 0 {
		fmt.Println("Positivo")
	} else if n < 0 {
		fmt.Println("Negativo")
	} else {
		fmt.Println("Nulo")
	}
}
