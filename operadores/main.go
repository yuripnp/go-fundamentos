package main

import "fmt"

func main() {
	// OPERADORES MATEMATICOS
	soma := 1 + 2
	subtracao := 4 - 3
	divisao := 6 / 2
	multiplicacao := 7 * 8
	restoDivisao := 9 % 2
	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	// OPERADORES DE ATRIBUICAO
	variavelUm := "string"
	fmt.Println(variavelUm)
	const variavelDois string = "string dois"
	variavelUm = "nova string"
	fmt.Println(variavelUm)

	// OPERADORES RELACIONAIS retorna booleano
	fmt.Println(1 > 2)  // maior
	fmt.Println(1 < 2)  // menor
	fmt.Println(1 >= 2) //maior ou igual
	fmt.Println(1 <= 2) // menor ou igual
	fmt.Println(1 == 2) // é igual
	fmt.Println(1 != 2) // é diferente

	// OPERADORES LOGICOS
	verdadeiroUm, falsoUm, verdadeiroDois := true, false, true
	fmt.Println(verdadeiroUm && falsoUm)
	fmt.Println(verdadeiroDois || falsoUm)
	fmt.Println(!verdadeiroDois)

	// OPERADORES UNARIOS que interagem com apenas uma variavel
	numero := 10
	fmt.Println(numero)
	numero++
	fmt.Println(numero)
	numero += 15
	fmt.Println(numero)
	numero--
	fmt.Println(numero)
	numero -= 10
	fmt.Println(numero)
	numero *= 3
	fmt.Println(numero)
	numero /= 2
	fmt.Println(numero)
	numero %= 2
	fmt.Println(numero)

}
