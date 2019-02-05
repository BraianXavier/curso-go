package main

import "fmt"

func main() {
	// var aprovados map[int]string
	//mapas deven ser inicializados
	aprovados := make(map[int]string)

	aprovados[12124444364] = "MARIA"
	aprovados[32456789766] = "Pedro"
	aprovados[21345554343] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)

	}
	fmt.Println(aprovados[21345554343])
	delete(aprovados, 21345554343)
	fmt.Println(aprovados[21345554343]) //não mostrou nada pois a chave foi deletada na linha 20
}
