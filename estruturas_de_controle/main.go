package main

import "fmt"

func main() {
	// ---- estruturas de controle

	numero := 10
	if numero >= 15 {
		fmt.Println("Maior ou igual a 15")
	} else {
		fmt.Println("Menor que 15")
	}

	if outroNumero := numero; outroNumero > 0 { // ifinit a variavel se limita ao escopo
		fmt.Println("Numero maior que zero")
	}
}
