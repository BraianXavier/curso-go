package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	//fmt.Println(x / float64(y))
	fmt.Println(int(x) / y)

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//cuidado..
	fmt.Println("Teste" + string(123))
	/*nessa forma você estará convertendo o valor 123
	da tabela ask assim como se voce fizer o seguinte bloco de código
	fmt.Println("Teste" + string(97))você receberá como resposta
	a letra "a" que é equivalente ao numero 97 na tabela ask*/

	//int para string
	fmt.Println("Teste" + strconv.Itoa(123))

	//string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}
