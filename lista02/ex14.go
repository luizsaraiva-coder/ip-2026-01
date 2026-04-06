// Desenvolva um programa para calcular e imprimir o preço final de um carro. O valor do preço inicial de
// fábrica é fornecido pelo usuário. O carro pode ter as seguintes opções:
// a) Ar condicionado (R$ 1750,00)
// b) Pintura metálica (R$ 800,00)
// c) Vidro elétrico (R$ 1200,00)
// d) Direção hidráulica (R$ 2000,00)

package main

import f "fmt"

func main() {

	var valorInicial, valorFinal float64
	var ac, pm, ve, dh int

	f.Println("Digite o preço de fábrica do carro: ")
	f.Scan(&valorInicial)

	valorFinal = valorInicial

	f.Println("Deseja ar condicionado? (1-Sim / 2-Não): ")
	f.Scan(&ac)
	if ac == 1 {
		valorFinal += 1750
	}

	f.Println("Deseja pintura metálica? (1-Sim / 2-Não): ")
	f.Scan(&pm)
	if pm == 1 {
		valorFinal += 800
	}

	f.Println("Deseja vidro elétrico? (1-Sim / 2-Não): ")
	f.Scan(&ve)
	if ve == 1 {
		valorFinal += 1200
	}

	f.Println("Deseja direção hidráulica? (1-Sim / 2-Não): ")
	f.Scan(&dh)
	if dh == 1 {
		valorFinal += 2000
	}

	f.Printf("Preço final do carro: R$ %.2f\n", valorFinal)

}
