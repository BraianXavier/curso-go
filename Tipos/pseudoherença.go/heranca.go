package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}
type ferrari struct {
	carro       //campos anônimos
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "f48"
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Printf("A ferrari %s esta com turbo ligado? %v\n", f.nome, f.turboLigado)
}
