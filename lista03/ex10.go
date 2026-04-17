// Faça um programa que monte a sequência de Fibonacci com N termos.
// A sequência de Fibonacci é dada por: 0 – 1 – 1 – 2 – 3 – 5 – 8 – 13 – 21 – 34 – 55 – ...
// O primeiro e o segundo termos da sequência de Fibonacci são 0 e 1.
// Considere que o usuário irá informar um número N >= 3.

package main

import f "fmt"

func main() {

	var n int

	f.Print("Digite a quantidade de termos (N >= 3): ")
	f.Scan(&n)

	a := 0
	b := 1

	f.Print(a, " ", b, " ")

	for i := 3; i <= n; i++ {
		c := a + b
		f.Print(c, " ")

		a = b
		b = c
	}
}