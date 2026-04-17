// Escreva um programa que receba vários números, calcule e mostre:
// a) a soma dos números digitados;
// b) a quantidade de números digitados;
// c) a média dos números digitados;
// d) o maior número digitado;
// e) o menor número digitado;
// f) a média dos números pares;
// g) a percentagem dos números ímpares entre todos os números digitados.
// Finalize a entrada de dados com a digitação do número 30.000.

package main

import f "fmt"

func main() {

	var num int
	var soma, qtd, maior, menor, somaPares, qtdPares, qtdImpares int

	for {
		f.Print("Digite um número (30000 para sair): ")
		f.Scan(&num)

		if num == 30000 {
			break
		}

		if qtd == 0 {
			maior, menor = num, num
		}

		soma += num
		qtd++

		if num > maior {
			maior = num
		}
		if num < menor {
			menor = num
		}

		if num%2 == 0 {
			somaPares += num
			qtdPares++
		} else {
			qtdImpares++
		}
	}

	f.Println("\nSoma:", soma)
	f.Println("Quantidade:", qtd)

	if qtd > 0 {
		f.Println("Média:", float64(soma)/float64(qtd))
	}

	f.Println("Maior:", maior)
	f.Println("Menor:", menor)

	if qtdPares > 0 {
		f.Println("Média dos pares:", float64(somaPares)/float64(qtdPares))
	} else {
		f.Println("Não há números pares")
	}

	if qtd > 0 {
		f.Println("Percentual de ímpares:", float64(qtdImpares)/float64(qtd)*100, "%")
	}
}
