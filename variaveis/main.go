package main

import "fmt"

func main() {
	// variaveis fortemente tipadas e toda variavel tem que ser usada
	// implicido e explicido
	var variavel string = "variavel explicidamente tipada"
	fmt.Println(variavel)

	variavelDois := "variavel implicidamente tipada"
	fmt.Println(variavelDois)

	// atribuição em grupo

	var (
		variavelTres   string = "lalala"
		variavelQuatro int    = 12
	)
	fmt.Println(variavelTres, " ", variavelQuatro)

	variavelCinco, variavelSeis := "variavel novas", "variaveis em grupo"
	variavelsete, variaveloito := "texto var", 12
	fmt.Println(variavelCinco, " ", variavelSeis, variavelsete, variaveloito)

	const variavelfixa string = "imutavel" // mesmo que var mas não muda
	const variavelConstante int = 12
	const (
		teste1 int = 12
		teste2 int = 13
	)
	fmt.Println(variavelConstante)
	// inversão de variaveis
	variavelCinco, variavelSeis = variavelSeis, variavelCinco

}
