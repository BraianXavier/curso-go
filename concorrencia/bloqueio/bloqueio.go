package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 //operação bloqueante
	fmt.Println("Só depois de ser lido")
}

func main() {
	c := make(chan int) // canal sem buffer

	go rotina(c)

	fmt.Println(<-c) // operação bloqueamente
	fmt.Println("foi lido")
	fmt.Println(<-c) // deadlock
	fmt.Println("Fim")
}
