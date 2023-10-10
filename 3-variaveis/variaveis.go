package main

import "fmt"

func main() {
	//declaração de variaveis

	//explicita
	var variavel1 string = "Variavel 1"

	//implicita
	variavel2 := "Variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Olá!"
		variavel4 int    = 123
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel 5", "Variavel 6"

	fmt.Println(variavel5, variavel6)

	//contantes
	const constante1 = "Constante 1"
	fmt.Println(constante1)

	//swap de valores
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
