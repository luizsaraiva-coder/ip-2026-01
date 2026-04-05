// Escreva um programa que receba um número inteiro positivo de 3 casas e imprima o algarismo da casa das
// dezenas. Não se esqueça de testar para ver se o número informado tem realmente 3 casas.

package main

import f "fmt" 

func main() {

	var n int 

	f.Print("Digite um número de 3 casos")]
	f.Scan(&n)

	if n < 100 || n > 999 {
		fmt.Println("Número inválido. Deve ter 3 casas.")
		return
	}

	dezena := (num / 10) % 10

	fmt.Println("Algarismo da casa das dezenas:", dezena)
}
