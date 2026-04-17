// Escreva um programa que receba a idade, a altura e o peso de várias pessoas. Calcule e imprima:
// - a quantidade de pessoas com idade superior a 50 anos;
// - a média das alturas das pessoas com idade entre 10 e 20 anos;
// - a porcentagem de pessoas com peso inferior a 40 quilos entre todas as pessoas analisadas.
// Considere que os dados informados são válidos. Pergunte ao usuário se ele deseja continuar digitando dados ou
// não (Exemplo: 1 - Sim, Outro valor diferente de 1 - Não).

package main

import f "fmt"

func main() {

	var idade, opcao int
	var altura, peso float64

	mais50 := 0
	totalPessoas := 0
	somaAltura := 0.0
	qtdAltura := 0
	pesoMenor40 := 0

	for {
		f.Print("Digite a idade: ")
		f.Scan(&idade)

		f.Print("Digite a altura: ")
		f.Scan(&altura)

		f.Print("Digite o peso: ")
		f.Scan(&peso)

		totalPessoas++

		// idade > 50
		if idade > 50 {
			mais50++
		}

		// idade entre 10 e 20
		if idade >= 10 && idade <= 20 {
			somaAltura += altura
			qtdAltura++
		}

		// peso < 40
		if peso < 40 {
			pesoMenor40++
		}

		f.Print("Deseja continuar? (1 - Sim / outro valor - Não): ")
		f.Scan(&opcao)

		if opcao != 1 {
			break
		}
	}

	// Resultados
	f.Println("\nResultados:")

	f.Println("Quantidade de pessoas com mais de 50 anos:", mais50)

	if qtdAltura > 0 {
		media := somaAltura / qtdAltura
		f.Println("Média das alturas (10 a 20 anos):", media)
	} else {
		f.Println("Não há pessoas entre 10 e 20 anos.")
	}

	if totalPessoas > 0 {
		porcentagem := float64(pesoMenor40) / float64(totalPessoas) * 100
		f.Println("Porcentagem com peso < 40kg:", porcentagem, "%")
	}
}