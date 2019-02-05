package main

import "fmt"

func notaPraConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "b"
	case 6, 5:
		return "c"
	case 4, 3:
		return "d"
	case 2, 1, 0:
		return "e"
	default:
		return "Nota invalida"
	}
}

func main() {

	fmt.Println(notaPraConceito(10))
	fmt.Println(notaPraConceito(9))
	fmt.Println(notaPraConceito(8))
	fmt.Println(notaPraConceito(6))
	fmt.Println(notaPraConceito(4))
	fmt.Println(notaPraConceito(2))
	fmt.Println(notaP, raConceito(55))
}
