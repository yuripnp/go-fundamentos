package main

import (
	"fmt"
	"module/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Teste", "testando")
	auxiliar.Escrever()
	erroEmail := checkmail.ValidateFormat("123")
	if erroEmail != nil {
		fmt.Println(erroEmail)
	} else {
		fmt.Println("sucess email")
	}
}
