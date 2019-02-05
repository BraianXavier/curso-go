package main

import "fmt"

func obtervaloraprovado(numero int) int {
	defer fmt.Println("fim!!!")
	if numero >= 5000 {
		fmt.Println("valor alto..")
		return 5000
	} else {
		fmt.Println("valor baixo...")
		return numero
	}
}

func main() {
	fmt.Println(obtervaloraprovado(6000))
}
