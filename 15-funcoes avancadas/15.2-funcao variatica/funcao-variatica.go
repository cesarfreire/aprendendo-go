package main

import "fmt"

// aceita qualquer quantidade de parâmetros com o ...
// apenas 1 por função, deve ser o último
func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}

}

func main() {
	totalSoma := soma(1, 4, 6, 7, 8, 65, 3, 2333, 5, 6, 4)
	fmt.Println(totalSoma)

	escrever("Olá mundo!", 4, 7, 8)
}
