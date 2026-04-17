// Escreva um programa que receba vários números inteiros e
// verifique se eles são ou não quadrados perfeitos. O
// programa deve terminar quando o usuário informar
// um número menor ou igual a zero. Obs.: Um número é
// quadrado perfeito quando tem um número inteiro como
// raiz quadrada. Não é permitido usar o comando sqrt
// em sua solução.

package main 

import f "fmt" 

func main () {

	var num int

	for {
		f.Print("Digite um número inteiro (<= 0 para sair): ")
		f.Scan(&num)

		if num <= 0 {
			f.Println("Encerrando o programa")
			break
		} else {
			quadradoPerfeito := false

			for i := 0; i*i <= num; i++ {
				if i*i == num {
					quadradoPerfeito = true
					break
		}
	}
		if quadradoPerfeito {
			f.Println("É um quadrado perfeito")
		}else {
			f.Println("Não é um quadrado perfeito")
		}
}
	}
}
