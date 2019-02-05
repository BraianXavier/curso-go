package main

import "fmt"

func inc1(n int) {
	n++
}

// revisão um ponteiro é representado por um *

func inc2(n *int) {
	//revisão:*é usado para desreferenciar, ou seja,
	//ter acesso o valor no qual o ponteiro aponta

	*n++
}

func main() {
	n := 1
	inc1(n) // por valor
	fmt.Println(n)

	//revisão: & usado para obter o endereco da variavel
	inc2(&n) // por referencia
	fmt.Println(n)
}
