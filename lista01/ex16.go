// Reajuste salarial - Att16
// Fazer um algoritmo que calcule e imprima o salário reajustado de um funcionário de acordo com as
// seguintes regras:
// • Salário de até R$ 300,00, reajuste de 50%;
// • Salário maior que R$300,00 reajuste de 30%;
// Entrada
// A entrada conterá uma linha com o salário do funcionário.
// Saída
// A saída deve conter, numa linha com a frase: SALARIO COM REAJUSTE = x, onde x é um valor real
// com duas casas decimais e corresponde ao valor do salário reajustado. Logo em seguida ao valor de x, o
// programa devem imprimir o caractere de quebra de linha: ‘\n’.

package main

import (
	"fmt"
)

func main() {
	var salario, novoSalario float64

	fmt.Print("Informe o salário do funcionário: ")
	fmt.Scan(&salario)

	if salario <= 300 {
		novoSalario = salario * 1.5
	} else {
		novoSalario = salario * 1.3
	}

	fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", novoSalario)
}
