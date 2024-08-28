package main

import "fmt"

//---------- FUNÇÃO RETORNO NOMEADO ---------------
func calculoMatematico(n1, n2 int) (soma int, multiplicacao int) {
	soma = n1 + n2
	multiplicacao = n1 * n2
	return
}

func calculo(n1, n2 float32) (float32, float32) {
	divisao := n1 / n2
	subtracao := n1 - n2
	return divisao, subtracao
}

// ------------------------------------------------

// ----------- FUNÇÃO VARIATICA -------------------

func soma(numeros ...int) {
	fmt.Print(numeros)
}

func teste(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func nomecalc(texto string, numeros ...int) {
	for _, numero := range numeros {
		if texto == "par" {
			if numero%2 == 0 {
				fmt.Println(texto, numero)
			}
		} else if texto == "impar" {
			if numero%2 != 0 {
				fmt.Println(texto, numero)
			}
			fmt.Println(texto, numero)
		}
	}
}

//--------------------------------------------------

//----------- FUNÇÃO RECURSIVA ---------------------
// uma função que chama ela mesma, (toda função recursiva tem que ter parada, para não estourar a pilha)
func fibonati(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonati(posicao-2) + fibonati(posicao-1)
}

func main() {
	som, mult := calculoMatematico(4, 8)
	fmt.Println(som, mult)
	teste := teste(1, 2, 3, 4, 5)
	fmt.Println(teste)
	nomecalc("par", 1, 4, 5, 3, 2, 10)

	//-----------FUNÇÕES ANONIMAS-----------------------
	func() {
		fmt.Println("teste funcao anonima")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("passando aqui os parametros")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("teste retorno")

	fmt.Println(retorno)
	//--------------------------------------------------

	numerouint := fibonati(10)
	fmt.Println(numerouint)
}
