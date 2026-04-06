// Elabore um programa que calcule o valor a ser pago por um produto considerando o preço normal de etiqueta
// e a escolha da condição de pagamento. Utilize os códigos da tabela a seguir para saber qual a condição de
// pagamento escolhida e efetuar o cálculo adequado.
// Código                                                          Condição de Pagamento
// 1                                                À vista, dinheiro ou cheque, 10% de desconto
// 2                                                 À vista, cartão de crédito, 5% de desconto
// 3                                                   Em 2 vezes, preço normal de etiqueta sem juros
// 4                                                  Em 3 vezes, preço normal de etiqueta + 10% de juros

package main

import (
	"fmt"
)

func main() {
	var opcao int
	var precoB, precoF float64

	fmt.Println("Informe o valor do produto na etiqueta:")
	fmt.Scan(&precoB)

	fmt.Println("Informe a opção de pagamento:")
	fmt.Println("1 - À vista, dinheiro ou cheque")
	fmt.Println("2 - À vista, cartão de crédito")
	fmt.Println("3 - Parcelado 2 vezes")
	fmt.Println("4 - Parcelado 4 vezes")
	fmt.Scan(&opcao)

	precoF = precoB

	switch opcao {
	case 1:
		precoF *= 0.9
	case 2:
		precoF *= 0.95
	case 3:
		precoF = precoB
	case 4:
		precoF *= 1.1
	default:
		fmt.Println("Opção inválida!")
		return
	}

	fmt.Printf("O valor do pagamento total de acordo com a opção %d é R$ %.2f\n", opcao, precoF)
}
