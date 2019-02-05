package main

import "fmt"

func main() {
	funcionariosESalarios := map[string]float64{
		"josé joão":      11345.45,
		"Gabriela Silva": 156664.12,
		"Pedro Junior":   332444.23,
	}

	funcionariosESalarios["Rafael filho"] = 5555.33
	delete(funcionariosESalarios, "inexistente")

	for nome, salario := range funcionariosESalarios {
		fmt.Println(nome, salario)
	}
}
