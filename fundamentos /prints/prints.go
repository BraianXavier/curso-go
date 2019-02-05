package main

import "fmt"

func main() {
	fmt.Print("mesma")
	fmt.Print("Linha!") //bloco executado na mesma linha.

	fmt.Println("Linha")    // essa linha de comando contunuara na mesma linha, e dos dela é que ocorre a quebra.
	fmt.Println("de baixo") //esse bloco ja estará quebrado, ou seja, ja sera iniciado em outra linha.

	x := 3.141516

	//fmt.Println("O valor de x é " + x)
	xs := fmt.Sprint(x)
	fmt.Println(" O valor de x é " + xs)
	fmt.Println("o valor de x é ", x)
	//valores formatados!!!!!!!!!!!11
	fmt.Printf("O valorde x é %.2f.", x)

	a := 1
	b := 1.99999
	c := false
	d := "opa!"

	fmt.Printf("\n %d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n %v %v %v %v", a, b, c, d) //a configuração "%v" formata todas as variaveis.

}
