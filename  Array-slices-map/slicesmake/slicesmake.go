package main

import "fmt"

func main() {
	s := make([]int, 10) //criando um slice e um array interno ao mesmo tempo usando make
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20) // dizendo qual o tamanho do slices e do array usando o make
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) // adicionando dados em um slices apartir de um append
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1) //adicionando um valor a mais que a capacidade normal do array interno, o array se duplica automaticamente.
	fmt.Println(s, len(s), cap(s))
}
