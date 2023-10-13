package main

import "fmt"

// passando um parâmetro por cópia
func inverteSinal(numero int) int {
	return numero * -1
}

// recebe um ponteiro de int
// não precisa de retorno pois altera o valor direto na memória
func inverteSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	//é repassado uma cópia do numero, e não o número em si
	numeroInvertido := inverteSinal(numero)
	fmt.Println(numeroInvertido)

	//altera diretamente na variável
	novoNumero := 40
	fmt.Println(novoNumero)
	//passa o endereço de memória ao invés do valor
	inverteSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
