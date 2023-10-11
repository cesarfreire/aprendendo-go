package main

import "fmt"

func main() {

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	//if init - limita ao escopo do if
	if outroNumero := 10 + numero; outroNumero > 0 {
		fmt.Println("É maior que zero")
	} else if outroNumero < 10 {
		fmt.Print("Menor que 10")
	} else if outroNumero > 10 {
		fmt.Print("Maior que 10")
	} else {
		fmt.Print("fim")
	}

	//switch

}
