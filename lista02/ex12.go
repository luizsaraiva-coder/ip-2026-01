// Construa um programa que receba a idade de uma pessoa e classifique-a seguindo o critério apresentado a
// seguir. Considere que a idade é um valor inteiro e que será informada de forma válida.
// Idade                               Classificação
// 0 a 2 anos                          Recém-nascido
// 3 a 11 anos                           Criança
// 12 a 19 anos                         Adolescente
// 20 a 55 anos                            Adulto
// Acima de 55 anos                          Idoso

package main

import f "fmt"

func main() {

	var idade int

	f.Println("Digite a idade: ")
	f.Scan(&idade)

	if idade >= 0 && idade <= 2 {
		f.Print("Recém-nascido")

	} else if idade >= 3 && idade <= 11 {
		f.Print("Criança")

	} else if idade >= 12 && idade <= 19 {
		f.Print("Adolescente")

	} else if idade >= 20 && idade <= 55 {
		f.Print("Adulto")

	} else {
		f.Print("Idoso")
	}

}
