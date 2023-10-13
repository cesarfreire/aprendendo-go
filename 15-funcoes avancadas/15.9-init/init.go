package main

import "fmt"

var n int

// utilizada para fazer um "setup" antes da função main
func init() {
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Função main sendo excutada")
	fmt.Println(n)
}
