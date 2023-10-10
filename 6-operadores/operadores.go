package main

import "fmt"

func main() {
	// ARITMÉTICOS
	var numero1 int16 = 32
	var numero2 int16 = 258

	soma := numero1 + numero2

	fmt.Println(soma)

	//RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 != 2)
	fmt.Println(1 == 2)

	//LÓGICOS
	fmt.Println(true && false)

	fmt.Println(true || true)

	fmt.Println(!true)

	//UNÁRIOS
	numero := 10
	numero++
	fmt.Println(numero)
	numero += 5
	fmt.Println(numero)
	numero--
	fmt.Println(numero)

	//TERNÁRIOS
	//Não existe em go
	//texto := numero > 5 ? "Maior que 5" : "Menor que 5"
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)
}
