package main

import "fmt"

func closure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

func main() {
	x := 20
	fmt.Println(x)
	/*O closure respeita sua linhagem de código, por isso uma vez declarada
	  que esta sendo apontato o valor atribuido a ela mesmo que seja em um outro bloco de funcao*/
	imprimeX := closure()
	imprimeX()
}
