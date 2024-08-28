package main

import "fmt"

func soma(n1 int8, n2 int8) int8 {
	return n1 + n2
}
func somaSubtracao(n1, n2 int16) (int16, int16) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	somar, subtrair := somaSubtracao(100, 30)
	fmt.Println(somar, subtrair)
	soma := soma(1, 4)

	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("texo funcao anonima")
	fmt.Println(resultado)

	resultadoSoma, _ := somaSubtracao(12, 92) // quando temos duas saidas mas ignoramos uma
	fmt.Println(resultadoSoma)
}
