package main

import (
	"errors"
	"fmt"
)

func main() {
	var numeroUm int8 = 99 // -99 - 99
	fmt.Println(numeroUm)

	var numeroDois int16 = 9999 // -9.999 - 9.999
	fmt.Println(numeroDois)

	var numeroTres int32 = 999999999 // -999.999.999 - 999.999.999
	fmt.Println(numeroTres)

	var numeroQuatro uint8 = 99 // 0 - 99
	fmt.Println(numeroQuatro)

	//alias
	//int32 = rune
	var numeroCinco rune = 1234
	fmt.Println(numeroCinco)

	//BYTE
	//UINT8 = byte
	var numeroSeis byte = 99
	fmt.Println(numeroSeis)

	//numero ponto flutuante
	var (
		numeroSete float32 = 12.999
		numeroOito float64 = 12.999999
	)
	fmt.Println(numeroSete)
	fmt.Println(numeroOito)

	// string char
	var texto string = "colocar testo"
	fmt.Println(texto)

	char := 'A' // numero da posição na memoria
	fmt.Println(char)

	var numeroindef int32
	fmt.Println(numeroindef)

	var character string
	fmt.Println(character)

	// booleano
	var boleano bool
	fmt.Println(boleano)

	//error
	var erro error = errors.New("Teste de erro")
	fmt.Println(erro)
}
