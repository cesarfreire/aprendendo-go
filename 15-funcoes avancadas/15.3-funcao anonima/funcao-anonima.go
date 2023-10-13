package main

import "fmt"

func main() {
	// o () executa a funcao logo após criada
	func() {
		fmt.Println("Olá mundo!")
	}()

	// com parâmetros
	func(texto string) {
		fmt.Println(texto)
	}("Olá mundo 2, via parâmetro")

	// com retorno
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando parâmetro")

	fmt.Println(retorno)
}
