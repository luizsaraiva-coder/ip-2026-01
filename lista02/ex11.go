// Escreva um programa que receba o valor de x, calcule e imprima o valor de f(x), dado que:
// f(x) = 8 / (2-x)

package main

import f "fmt"

func main() {

	var x, resultado float64

	f.Println("Digite o valor de x: ")
	f.Scan(&x)

	if x == 2 {
		f.Println("Erro: divisão por zero (x não pode ser 2)")
		return
	}

	resultado = 8 / (2 - x)

	f.Printf("f(x) = %.2f\n", resultado)

}
