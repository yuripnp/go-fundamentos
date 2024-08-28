package main

import "fmt"

type pessoa struct {
	nome   string
	altura float32
	idade  int8
}
type estudante struct { // pegando todos os campos que está em pessoa para estudante, é como se fosse herança
	pessoa    // nesse caso você não precisa especificar o tipo
	curso     string
	faculdade string
}

type professor struct {
	pessoa  pessoa
	salario float32
	materia string
}

func main() {
	fmt.Println("Herança")
	p1 := pessoa{"Nome", 1.67, 34}
	e1 := estudante{p1, "engenharia", "faculdade x"}

	fmt.Println(e1.nome, e1.curso, e1.faculdade)

	p2 := pessoa{"Nome dois", 1.90, 29}
	pro1 := professor{p2, 3000.00, "matematica"}
	fmt.Println(pro1.pessoa.nome, pro1.salario, pro1.materia)

}
