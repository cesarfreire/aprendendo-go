package main

import "fmt"

func main() {

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++

	fmt.Println(variavel1, variavel2)

	//são copias

	//ponteiro é uma referência de memória
	var variavel3 int
	//*int > ponteiro de int
	var ponteiro *int

	variavel3 = 100
	//Endereco da memória &var
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)
	//desreferênciação *var
	fmt.Println(variavel3, *ponteiro)
}
