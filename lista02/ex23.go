// Criar um programa que leia a idade de uma pessoa e que mostre a sua classe eleitoral:
// - Não-eleitor (abaixo de 16 anos);
// - Eleitor obrigatório (entre 18 e 65 anos);
// - Eleitor facultativo (entre 16 e 18 anos e maior de 65 anos).

package main

import (
	"fmt"
)

func main() {
	var idade int

	fmt.Println("Informe a idade:")
	fmt.Scan(&idade)

	if idade < 0 {
		fmt.Println("Idade inválida!")
		return
	}

	switch {
	case idade < 16:
		fmt.Println("Não-eleitor")
	case idade < 18 || idade > 65:
		fmt.Println("Eleitor facultativo")
	default:
		fmt.Println("Eleitor obrigatório")
	}
}
