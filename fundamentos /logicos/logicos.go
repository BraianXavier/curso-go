package main

import "fmt"

//func main(){}

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprartv59 := trab1 && trab2
	comprartv39 := trab1 != trab2
	comprarsorvete := trab1 || trab2

	return comprartv59, comprartv39, comprarsorvete

}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("tv50: %t, tv32: %t , sorvete: %t , saudavel:%t", tv50, tv32, sorvete, !sorvete)
}
