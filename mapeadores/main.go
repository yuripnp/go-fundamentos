package main

import "fmt"

func main() {
	// as chaves tem sempre o mesmo type e os valores tbm sempre tem o mesmo type
	usuario := map[string]string{
		"nome":      "pedro",
		"sobrenome": "silva",
	}
	fmt.Println(usuario["nome"])
	fmt.Println(usuario)

	usuarioDois := map[string]map[string]string{
		"nome": {
			"primeiro": "segundo",
			"nome":     "sobrenome",
		},
	}

	fmt.Println(usuarioDois)
	delete(usuarioDois, "nome") // tirar o valor de um map, usamos a função map: primeiro colocamos o map e depois a chave
	fmt.Println(usuarioDois)

	// adicionando um valor
	usuarioDois["signo"] = map[string]string{
		"nome": "gercio",
	}
	fmt.Println(usuarioDois)
}
