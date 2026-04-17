// Faça um programa que imprima todos os números inteiros de 1 (inclusive) até 20 (inclusive) e também a soma
// de todos eles.

package main

import f "fmt"

func main() {

	soma := 0

	for i := 1; i <= 20; i++ {
		f.Println(i)
		soma += i
	}

	f.Println("Soma:", soma)
}
