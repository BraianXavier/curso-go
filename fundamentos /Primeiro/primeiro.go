package main /*todo programa "go" roda dentro de um pacote.
usado para quebrar um grande sistema em varias partes para facilitar a busca*/

import "fmt"

//a porta de entrada de um pacote é a função main.
func main() {
	fmt.Print("Primeiro ") //comando para mesma linha
	fmt.Print(" programa!")
	//o ponto e virgula não é necessário ao final de todo bloco, porque o go ja usa isso em modo auto.
}
