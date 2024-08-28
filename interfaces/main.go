package main

import (
	"fmt"
	"math"
)

// interface generica. Um método que aceita qualquer tipo
func generica(interfac interface{}) {
	fmt.Println(interfac)
}

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("Temos aqui a area da sua forma %.02f\n", f.area())
}

type triangulo struct {
	altura  float64
	largura float64
}

func (t triangulo) area() float64 {
	return t.altura * t.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	c := circulo{2.89}
	t := triangulo{altura: 1.20, largura: 2.90}
	escreverArea(c)
	escreverArea(t)
	generica("teste string")
	generica(true)
	generica(12)

	// sem controle da interface, pode-se virar bagunça, como no map
	mapa := map[interface{}]interface{}{
		1:        "string",
		"string": "string",
		true:     112.12,
	}

	fmt.Println(mapa)
}
