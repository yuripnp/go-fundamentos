package main

import "fmt"

func funcao01() {
	fmt.Println("imprime funcao 1")
}

func funcao02() {
	fmt.Println("imprime funcao 2")
}

func imprimirResultado(n1, n2 float32) bool {
	// vai ser executado por ultimo
	defer fmt.Println("Média calculada, o resultado sera retornado:")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	//defer funcao01()
	//funcao02()

	fmt.Println(imprimirResultado(8, 5))
}
