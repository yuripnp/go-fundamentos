package main

import (
	"fmt"
)

func main() {
	// Vamos tentar fazer uma divisão por zero para provocar um pânico
	safeDivision(4, 2)
	safeDivision(4, 0)
	fmt.Println("Program continues after recover.")
}

func safeDivision(a, b int) {
	// Usando defer para garantir que a função recover seja chamada
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Se b for 0, o programa entrará em pânico
	if b == 0 {
		panic("division by zero")
	}

	fmt.Println("Result:", a/b)
}
