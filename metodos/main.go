package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	email string
}

func (u usuario) salvar() {
	fmt.Println("Dentro do método salvar")
	fmt.Printf("salvando o usuario %s no banco de dados\n", u.nome)
}

// com ponteiro ele persiste
func (u *usuario) fezAniversario() {
	u.idade++
}

func (u usuario) maiorDeIdade(idadeMinima uint8) bool {
	if u.idade >= idadeMinima {
		fmt.Printf("você tem %d anos de idade, pode passar\n", u.idade)
		return true
	} else {
		fmt.Printf("Não pode entrar você tem %d anos\n", u.idade)
		return false
	}
}

func main() {
	usuario01 := usuario{"Lucas", 16, "usuario@email.com"}
	fmt.Println(usuario01)
	usuario01.salvar()
	usuario01.maiorDeIdade(18)
	usuario01.fezAniversario()
	usuario01.fezAniversario()
	usuario01.maiorDeIdade(18)
}
