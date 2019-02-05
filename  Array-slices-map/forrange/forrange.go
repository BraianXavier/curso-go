package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} //compilador conta!

	for i, numeros := range numeros {
		fmt.Printf("%d) %d\n", i+1, numeros)
	}

	for _, numeros := range numeros {
		fmt.Println(numeros)
	}
}
