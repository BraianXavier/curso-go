package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

//Método: funcao com reciver (receptor)

func (p produto) produtoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "lapis",
		preco:    12.90,
		desconto: 0.20,
	}

	fmt.Println(produto1, produto1.produtoComDesconto())

	//forma reduzida de se escrever um struct

	produto2 := produto{"notebook", 1989.90, 0.10}
	fmt.Println(produto2.produtoComDesconto())

}
