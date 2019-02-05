package main

import "fmt"

//meu teste
/*func main() {
	a1 := [6]int{1, 2, 3, 4, 5, 6}

	s1 := a1[4]

	s2 := a1[3]

	fmt.Println(a1, s1, s2)
}*/

func main() {
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)
	fmt.Println(s1, s2)

	//se adicionar um numero ao s1 ele ira mudar tando o s1 quanto o s2 pois ambos apontam para o mesmo array interno
	s1[0] = 7
	fmt.Println(s1, s2)
}
