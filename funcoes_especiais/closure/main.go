package main

import "fmt"

func clousure() func() {
	texto := "dentro da funcão clousure"
	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}
func main() {
	texto := "dentro da minha funcao main"
	fmt.Println(texto)
	funcaoNova := clousure()
	funcaoNova()
}
