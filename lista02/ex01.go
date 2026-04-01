// 1. Escreva um programa que leia um valor inteiro e diga se o número informado é par ou ímpar.

package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Informe um número: ")
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println("Par")
	} else {
		fmt.Println("Ímpar")
	}
}
