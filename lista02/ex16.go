// Desenvolver um programa que leia os coeficientes (A, B e C) de uma equação do segundo grau ( Ax² + Bx + C
// =0) e que calcule suas raízes. O programa deve mostrar, quando possível, o valor das raízes calculadas e a
// classificação das mesmas: “RAÍZES IMAGINÁRIAS”, “RAIZ ÚNICA” ou “RAÍZES DISTINTAS”.

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	delta, rDelta, r1, r2, rUnica := 0.0, 0.0, 0.0, 0.0, 0.0

	fmt.Println("Informe o valores de A, B e C:")
	fmt.Scan(&a, &b, &c)

	if a == 0 {
		fmt.Println("O valor de A não pode ser zero em uma equação do 2º grau.")
	}

	delta = math.Pow(b, 2) - (4 * a * c)

	switch {
	case delta > 0:
		rDelta = math.Sqrt(delta)
		r1 = ((-b) + rDelta) / (2 * a)
		r2 = ((-b) - rDelta) / (2 * a)
		fmt.Printf("Delta positivo\nValores de x1, x2 são %.3f e %.3f\n", r1, r2)
	case delta == 0:
		rUnica = (-b) / (2 * a)
		fmt.Printf("Delta igual a zero\nValor da raiz é %.3f\n", rUnica)
	case delta < 0:
		fmt.Println("Raízes imaginárias!\n")

	}

}
