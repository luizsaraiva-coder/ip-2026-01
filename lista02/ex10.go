// Escreva um programa que leia o destino de um passageiro e se a viagem inclui retorno (ou seja, se é
// só ida, ou se é ida e volta). Informe o preço da passagem de acordo com a tabela a seguir.
// Destino                                 Ida                                        Ida e Volta
// 1 - Região Norte                      R$ 500,00                                    R$ 900,00
// 2 - Região Nordeste                   R$ 350,00                                    R$ 650,00
// 3 - Região Centro-Oeste               R$350,00                                    R$ 600,00
// 4 - Região Sul                        R$ 300,00                                   R$ 550  ,00
// Sugestão: Considere “1” representando a Região Norte, “2” para Nordeste, “3” para Centro-Oeste e
// “4” para Sul. Para ver se a viagem inclui retorno, considere “1”- para sim (ida e volta) e “2” - para
// não (só ida).
// Cuidado com valores inválidos tanto para o destino quanto para o fato de incluir ou não o retorno.

package main

import f "fmt"

func main() {

	var destino, retorno int
	var preço float64

	f.Println("Diga o destino(1-Norte, 2-Nordeste, 3-Centro-Oeste, 4-Sul): ")
	f.Scan(&destino)

	f.Println("Deseja retorno(1-Sim, 2-Não)? ")
	f.Scan(&retorno)

	if (destino < 1 || destino > 4) || (retorno != 1 && retorno != 2) {
		f.Println("Valor inválido")
	}

	switch destino {
	case 1:
		if retorno == 1 {
			preço = 900.00
		} else {
			preço = 500.00
		}
	case 2:
		if retorno == 1 {
			preço = 650.00
		} else {
			preço = 350.00
		}
	case 3:
		if retorno == 1 {
			preço = 600.00
		} else {
			preço = 350.00
		}
	case 4:
		if retorno == 1 {
			preço = 550.00
		} else {
			preço = 300.00
		}
	default:
		f.Println("Erro inesperado no destino!")
		return
	}

	f.Printf("O valor da passagem é R$ %.2f\n", preço)
}
