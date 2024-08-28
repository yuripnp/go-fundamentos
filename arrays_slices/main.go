package main

import (
	"fmt"
	"reflect"
)

func main() {
	// array
	var arreiUm [5]string //vai ter 5 posições com strings
	fmt.Println(arreiUm)
	arreiUm[0] = "posicao inicial"
	fmt.Println(arreiUm)

	arreiDois := [5]int{1, 2, 3, 4, 5} // tamanho fixo
	fmt.Println(arreiDois)

	arreiTres := [...]int{1, 2, 3, 4, 5, 6, 7} // tamanho dinamico
	fmt.Println(arreiTres)

	// Slice
	slice := []int{10, 20, 33, 12, 10} // tamanho indeterminado são tipos diferentes
	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(arreiTres))

	// forma padrão
	slice = append(slice, 19)
	fmt.Println(slice)

	// desse jeito funciona como ponteiro
	sliceDois := arreiDois[1:3] // pega da posicao 1 até a posição 2 e joga no slice º_*
	fmt.Println(sliceDois)
	arreiDois[1] = 7
	fmt.Println(sliceDois)

	// Arrays internos
	fmt.Println("-------- Arrays Internos --------")
	sliceTres := make([]float32, 2, 3) // quero armazenar o : tipo, numero de posicoes, e capacidade maxima (aqui é feito um array interno)
	fmt.Println(sliceTres)
	//----------- tamanho
	fmt.Println(len(sliceTres)) // tamanho do slice
	//----------- capacidade
	fmt.Println(cap(sliceTres)) // capacidade maxima dele

	sliceTres = append(sliceTres, 4)
	sliceTres = append(sliceTres, 7)
	fmt.Println(sliceTres)
	fmt.Println(len(sliceTres))
	fmt.Println(cap(sliceTres)) // ele cria um array de forma interna, e se passa desse limite, ele dobra o limite.

	sliceQuatro := make([]float32, 5)
	fmt.Println(sliceQuatro)
	fmt.Println(cap(sliceQuatro))
}
