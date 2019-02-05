package main

func main() {
	i := 1

	//go não tem aritmética de ponteiros

	var p *int = nil

	p = &i // pegando o endreço da variável

	*p++ // desreferenciando (pegando o valor)

	i++

	//fmt.Println(p, +p, i, &i)

}
