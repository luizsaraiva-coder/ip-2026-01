// Um comerciante calcula o valor da venda, tendo em vista a tabela a seguir:
// Valor da Compra Valor da Venda
// Valor < R$ 10,00                          Lucro de 70%
// R$ 10,00 <= Valor < R$ 30,00              Lucro de 50%
// R$ 30,00 <= Valor < R$ 50,00              Lucro de 40%
// Valor >= 50,00                            Lucro de 30%
// Escreva um programa que leia o valor da compra e imprima o valor da venda. Cuidado com valor
// inválido de compra!

package main

import f "fmt"

func main() {

	var compra, venda float64

	f.Println("Digite o valor da compra")
	f.Scan(&compra)

	if compra <= 0 {
		f.Println("Valor de compra inválido")
		return
	}

	if compra < 10 {
		venda = compra * 1.70

	} else if compra < 30 {
		venda = compra * 1.50

	} else if compra < 50 {
		venda = compra * 1.40

	} else {
		venda = compra * 1.30
	}

	f.Printf("Valor da venda: R$ %.2f\n", venda)

}
