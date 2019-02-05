package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabi nascimento": 3002.44,
			"Guga Pereira":    7003.76,
		},
		"J": {
			"jonny": 8000.9,
		},
		"P": {
			"Pedro Junior": 1200.3,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
