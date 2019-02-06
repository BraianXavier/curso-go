package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)

		fmt.Printf("%s: %s(iteração %d)\n", pessoa, texto, i+1)

	}

}

func main() {
	//fale("Maria", "pq voce não fala cmg?", 3)
	//fale("joão", "só posso falar dps de voce", 1)

	fale("maria", "ei", 500)
	fale("joão", "ei", 500)
}
