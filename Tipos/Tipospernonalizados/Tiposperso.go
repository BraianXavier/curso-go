package main

import "fmt"

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim

}

func notaParaConceito(n nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	} else if n.entre(7.0, 8.99) {
		return "B"
	} else if n.entre(5.0, 7.99) {
		return "C"
	} else if n.entre(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(4.0))
	fmt.Println(notaParaConceito(4.0))
	fmt.Println(notaParaConceito(4.0))
	fmt.Println(notaParaConceito(4.0))
	fmt.Println(notaParaConceito(4.0))
	fmt.Println(notaParaConceito(4.0))
	fmt.Println(notaParaConceito(2.0))
	fmt.Println(notaParaConceito(2.0))
	fmt.Println(notaParaConceito(2.0))
	fmt.Println(notaParaConceito(3.0))
	fmt.Println(notaParaConceito(6.0))
	fmt.Println(notaParaConceito(4.0))
	fmt.Println(notaParaConceito(3.0))
	fmt.Println(notaParaConceito(4.0))
	fmt.Println(notaParaConceito(0.0))
	fmt.Println(notaParaConceito(1.0))
	fmt.Println(notaParaConceito(2.0))
	fmt.Println(notaParaConceito(9.0))
	fmt.Println(notaParaConceito(7.0))
	fmt.Println(notaParaConceito(6.0))

}
