package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//numeros inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("literal inteiro é ", reflect.TypeOf(32000))

	//sem sinal (só positivos )
	var b byte = 255
	fmt.Println("O byte é ", reflect.TypeOf(b))

	// com sinal.. int8 int16 int32 int64

	il := math.MaxInt64
	fmt.Println("o valor minimo de int é ", il)
	fmt.Println("O tipo de il é ", reflect.TypeOf(il))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicade (int32)
	fmt.Println("o rune é ", reflect.TypeOf(i2))
	fmt.Println(i2)

	//numeros reais float32 float64
	var x float32 = 49.99
	fmt.Println("O tipo de x é ", reflect.TypeOf(x))
	fmt.Println("O tipo do literal de 49.99 é ", reflect.TypeOf(49.99)) //float64

	//bolean
	bo := true
	fmt.Println("O tipo de bo é ", reflect.TypeOf(bo))
	fmt.Println(!bo) /// "!"usado para mostrar o valor contrário do que foi declarado em uma variavel booleana

	//string
	s1 := "Ola meu nome é braian"
	fmt.Println(s1 + "!")
	fmt.Println("o tamanho da string é ", len(s1))

	//String com multiplas linhas
	s2 := `Olá 
	meu 
	nome 
	é 
	leo`
	fmt.Println("O tamanho da string é ", len(s2))
	fmt.Println(s2)

	//char??
	//var x char = 'b' não existe o tipo "char"
	char := "a"
	fmt.Println("O tipo de char é ", reflect.TypeOf(char)) //nesse caso ele vai ser do tipo int32 e receberá o valor de uma var de String.
	fmt.Println(char)

}
