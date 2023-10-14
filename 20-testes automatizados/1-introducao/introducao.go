package main

import (
	"fmt"

	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Praça da Sé")

	fmt.Println(tipoEndereco)
}
