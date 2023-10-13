package main

import "fmt"

func closure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	//o go mantém a referência inicial da variavel texto
	funcaoNova := closure()
	funcaoNova()
}
