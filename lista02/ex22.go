// Desenvolver um programa que calcule o salário bruto e o salário líquido de um funcionário.
// • Dados de Entrada: Matrícula do funcionário (int);
// Quantidade de horas-extras trabalhadas.
// • Constantes: Salário Mínimo = R$ 788.00;
// Valor da Hora-Extra = R$ 10.00.
// Sabe-se:
// • Salário hora-extra = horas-extras * Valor da Hora-Extra;
// • Salário bruto = 3 * Salário Mínimo + Salário hora-extra;
// • Desconto INSS = 12 % do salário bruto, se salário bruto for maior que R$ 1500,00;
// • Desconto do Imposto de Renda = 20 % do Salário Bruto, se o mesmo for maior que R$ 2000,00;
// • Salário líquido = salário bruto – deduções.

package main

import (
	"fmt"
)

func main() {
	var matricula int
	var qtdHorasExtras, salarioExtra, salarioBruto, salarioLiqui float64

	const (
		salarioMinimo, valorHoraExtra = 788.00, 10.0
	)

	fmt.Println("Informe a matrícula do funcionário:")
	fmt.Scan(&matricula)
	fmt.Println("Informe a quantidade de horas extras trabalhadas:")
	fmt.Scan(&qtdHorasExtras)

	salarioExtra = valorHoraExtra * qtdHorasExtras
	salarioBruto = 3*salarioMinimo + salarioExtra

	salarioLiqui = salarioBruto

	if salarioBruto < 2000 {
		salarioLiqui *= 0.88
	} else {
		salarioLiqui *= 0.80
	}

	fmt.Printf("Funcionário %d\nSalário Bruto: R$%.2f\nSalário liquído: R$%.2f\n", matricula, salarioBruto, salarioLiqui)

}
