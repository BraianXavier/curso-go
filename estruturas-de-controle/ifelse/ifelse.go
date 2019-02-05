package main

import "fmt"

//Estrutura (se/senão)
func imprimirresultado(nota float64) {
	if nota >= 7 {
		fmt.Println("aprovado com a nota", nota)

	} else {
		fmt.Println("Reprovado com nota ", nota)
	}
}

func main() {
	imprimirresultado(7.3)
	imprimirresultado(5.1)
}
