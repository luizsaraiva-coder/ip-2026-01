// Faça um programa que leia uma data (dia, mês e ano, em uma variável inteira cada), e escreva a mesma data
// no formato dia de (mês por extenso) de ano.

package main

import f "fmt"

func main() {

	var dia, mes, ano int

	f.Print("Digite dia, mês e ano: ")
	f.Scan(&dia, &mes, &ano)

	meses := []string{"", "janeiro", "fevereiro", "março", "abril", "maio", "junho",
		"julho", "agosto", "setembro", "outubro", "novembro", "dezembro"}

	if mes < 1 || mes > 12 {
		f.Println("Mês inválido")
		return
	}

	f.Printf("%d de %s de %d\n", dia, meses[mes], ano)
}
