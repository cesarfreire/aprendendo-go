package main

import "fmt"

func funcao1() {
	fmt.Println("Executando funcao 1...")
}

func funcao2() {
	fmt.Println("Executando funcao 2...")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada, o resultado será retornado.")
	//é executado antes do return
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}

func main() {
	// defer - adiar
	// adia a execução da função 1 até o último momento possível
	// defer funcao1()
	// funcao2()

	fmt.Println(alunoEstaAprovado(7, 8))
}
