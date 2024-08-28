package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	email    string
	endereco endereco
}

type endereco struct {
	rua    string
	cidade string
	numero string
}

func main() {
	var user usuario
	fmt.Println(user)

	user.nome = "Yuri"
	user.idade = 30
	user.email = "yuri@gmail.com"

	fmt.Println(user)

	enderecoUsuarioNew := endereco{"passagem P4", "TO", "90A"}
	usuarioNew := usuario{"Bento", 23, "bento@hotmail.com", enderecoUsuarioNew}
	fmt.Println(usuarioNew)

	usuarioDois := usuario{nome: "chico"}
	fmt.Println(usuarioDois)
}
