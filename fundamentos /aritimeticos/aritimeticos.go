package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Módulo =>", a%b)

	//bitwise
	fmt.Println("E =>", a&b)   // 11 & 10 = 10
	fmt.Println("ou =>", a|b)  // 11 | 10 = 11
	fmt.Println("xor =>", a^b) //11 ^ 10 = 01

	c := 3.0
	d := 2.0

	//outras operaçoes usando math..
	fmt.Println("maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Exponenciação =>", math.Pow(c, d))

}
