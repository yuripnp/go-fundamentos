package main

import "fmt"

func main() {
	var variavel int16 = 10
	var variavelDois int16 = variavel
	fmt.Println(variavelDois, variavel)

	variavel++
	fmt.Println(variavelDois, variavel)

	// TRABALHANDO COM PONTEIRO
	var variavelUm int // pode ser receber um valor int
	var ponteiro *int  // recebe um endereço de uma variavel que é do tipo inteiro
	var variavelString string
	fmt.Println(variavelUm, ponteiro, variavelString)

	variavelUm = 30
	variavelString = "teste"
	ponteiro = &variavelUm // usamos o & para dizer que, queremos o apontamento daquela variavel sem isso dá erro
	// não pode (ponteiro = &variavelString) e nem assim (ponteiro = &variavelUm )
	fmt.Println(variavelUm, ponteiro, variavelString)
	fmt.Println(*ponteiro) // desreferenciação pegamos o valor daquela referencia
}
