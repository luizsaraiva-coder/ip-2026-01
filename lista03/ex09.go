// Faça um programa que receba 2 notas de N alunos. Calcule e mostre:
// a) a média aritmética das 2 notas de cada aluno;
// b) uma mensagem de acordo com as regras a seguir:

// Média Aritmética
// Até 3
// Entre 3 e 7
// De 7 para cima

// Mensagem
// Reprovado
// Exame
// Aprovado

// c) o total de alunos aprovados;
// d) o total de alunos de exame;
// e) o total de alunos reprovados;
// f) a média da classe.

package main

import f "fmt"

func main() {

	var n int
	var nota1, nota2 float64
	var aprovados, exame, reprovados, somaMedias int 

	f.Print("Digite a quantidade de alunos: ")
	f.Scan(&n)

	for i := 1; i <= n; i++ {

		f.Println("\nAluno", i)

		f.Print("Digite a primeira nota: ")
		f.Scan(&nota1)

		f.Print("Digite a segunda nota: ")
		f.Scan(&nota2)

		media := (nota1 + nota2) / 2
		somaMedias += media

		f.Println("Média:", media)

		if media <= 3 {
			f.Println("Reprovado")
			reprovados++
		} else if media < 7 {
			f.Println("Exame")
			exame++
		} else {
			f.Println("Aprovado")
			aprovados++
		}
	}

	f.Println("\nResultados finais:")
	f.Println("Aprovados:", aprovados)
	f.Println("Exame:", exame)
	f.Println("Reprovados:", reprovados)

	if n > 0 {
		f.Println("Média da classe:", somaMedias/float64(n))
	}
}