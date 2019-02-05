package main

import (
	"fmt"
	m "math"
)

//para definir uma constante deve serguir os seguintes passos "constante+nome da constante+ tipo da constante

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //tipo (float64)

	//forma reduzida de criar uma var
	area := PI * m.Pow(raio, 2)
	fmt.Println("a area da circunferencia é", area)
	//blocos de constantes e variaveis.
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	e, f := true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}
