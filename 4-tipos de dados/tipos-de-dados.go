package main

import (
	"errors"
	"fmt"
)

func main() {
	//INICIAL:
	//int = 0
	//string = ""
	//bool = false
	//error = nil

	//int8, int16, int32, int64

	//erro
	//var numero int8 = 1000

	var numero1 int64 = 1000000000000000000
	fmt.Println(numero1)

	numero2 := 565656
	fmt.Println(numero2)

	//uint
	//unsigned int
	var numero3 uint32 = 5555555
	fmt.Println(numero3)

	//alias
	//int32 = rune
	var numero4 rune = 123456
	fmt.Println(numero4)

	//uint8 = byte
	var numero5 byte = 123
	fmt.Println(numero5)

	//float32, float64
	//utiliza PONTO
	var numero6 float32 = 123.45
	fmt.Println(numero6)

	numero7 := 6578363.2
	fmt.Println(numero7)

	// STRINGS
	var texto string = "Olá!"
	fmt.Println(texto)

	texto2 := "Olá 2!"
	fmt.Println(texto2)

	//char
	char := 'B'
	fmt.Println(char)
	//retorna o número na tabela ASC

	//boolean
	var booleano1 bool = true
	fmt.Println(booleano1)

	//error
	var erro error
	fmt.Println(erro)
	//default é nil

	//error
	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)

}
