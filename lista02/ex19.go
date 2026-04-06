// Desenvolver um programa com as opções de calcular e imprimir o volume e a área da superfície de um cone
// reto, de um cilindro ou de uma esfera. O programa deverá ler a opção da figura desejada (1-cone / 2-cilindro /
// 3-esfera) e de acordo com a opção escolhida calcular e escrever o volume e a área da superfície da figura
// pedida. Fórmulas:
// Cone Reto:
// Volume= (Π∗raio²∗altura) / 3         Area=Π∗raio∗√(raio2+ altura2)
// Cilindro:
// Volume= (Π∗raio²∗altura)             Area=2∗Π∗raio∗altura
// Esfera:
// Volume=( 4/3 )∗Π∗raio³               Area=4∗Π∗raio²

package main

import (
	"fmt"
	"math"
)

func main() {

	var r, h, v, a float64
	var forma int

	fmt.Println("Escolha a forma geometrica para calcular área: \n1 - Cone Reto\n2 - Cilindro\n3 - Esfera.")
	fmt.Scan(&forma)

	fmt.Println("Informe, caso possua, o valor do raio e altura da figura:")
	fmt.Scan(&r, &h)

	switch forma {
	case 1:
		v = (math.Pi * math.Pow(r, 2) * a) / 3
		a = (math.Pi * r * math.Sqrt((math.Pow(r, 2) + math.Pow(a, 2))))
	case 2:
		v = (math.Pi * math.Pow(r, 2) * a)
		a = 2 * math.Pi * r * a

	case 3:
		v = (4 / 3.0) * math.Pi * math.Pow(r, 3)
		a = 4 * math.Pi * math.Pow(r, 2)
	}

	fmt.Printf("Forma número %d\nVolume: %.2f m³\nÁrea: %.2f m²\n", forma, v, a)

}
