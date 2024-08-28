package main

import (
	"fmt"
)

func main() {
	variavel := 0

	for variavel < 10 {
		//fmt.Println(variavel)
		variavel++
		//time.Sleep(time.Second)
	}
	fmt.Println(variavel)

	for j := 0; j < 10; j++ {
		fmt.Println(j)
		//time.Sleep(time.Second)
	}

	nomes := [3]string{"joão", "Pedro", "lucas"}

	for indice, nome := range nomes { // primeiro é indice e depois o valor
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	var palavrasComA []string
	nomesAleatorios := [5]string{"lucas", "pedro", "yuri", "dennis", "paulo"}
	for _, nome := range nomesAleatorios {
		for _, letra := range nome {

			if string(letra) == "a" {
				palavrasComA = append(palavrasComA, nome)
			}
		}
	}
	fmt.Println(palavrasComA)

	// percorrendo com o map
	usuarios := map[string]string{
		"001": "Pedro",
		"002": "Gabriel",
		"003": "Ricardo",
	}
	for chave, valor := range usuarios { //usando range em map
		fmt.Println(chave, valor)
	}

}
