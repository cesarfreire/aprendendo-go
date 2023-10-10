package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Matias"
	u.idade = 20
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos bobos", 0}
	usuario2 := usuario{"Jorge", 25, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Teste"}
	fmt.Println(usuario3)
}
