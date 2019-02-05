package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d", i)
	}
	fmt.Println("\nMisturando...")
	for i := 0; i <= 10; i++ {
		if 1%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("impar")
		}
	}

	for {
		//laço infinito

		fmt.Println("infinito")
		time.Sleep(time.Second)
	}
	// Veremos o foreach no capitulo de array
}
