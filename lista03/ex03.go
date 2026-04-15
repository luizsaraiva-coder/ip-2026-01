// Faça um programa que receba o salário de um funcionário chamado Carlos. Sabe-se que outro funcionário,
// João, tem salário equivalente a um terço do salário de Carlos. Carlos aplicará seu salário integralmente na
// caderneta de poupança, que está rendendo 2% ao mês, e João aplicará seu salário integralmente no fundo de
// renda fixa, que está rendendo 5% ao mês. O programa deverá calcular e mostrar a quantidade de meses
// necessários para que o valor pertencente a João iguale ou ultrapasse o valor pertencente a Carlos.

package main

import f "fmt"

func main() {

	var salarioCarlos float64

	// Informe o valor
	f.Println("Informe o salário de Carlos: ")
	f.Scan(&salarioCarlos)

	salarioJoao := salarioCarlos / 3

	meses := 0

	// Laço de repetição
	for salarioJoao <= salarioCarlos {
		salarioCarlos += salarioCarlos * 0.02
		salarioJoao += salarioJoao * 0.05
		meses++
	}

	// Exibição do resultado
	f.Printf("Quantidade de meses necessários: %d\n", meses)
}
