package main

import (
	"fmt"
	"time"
)

func main() {
	// concorrencia != paralelismo
	// executa mas não espera a finalização para seguir o programa
	go escrever("Olá mundo!") //goroutine
	escrever("Programando em GO!")

	//não funciona adicionar go em tudo
	// go escrever("Olá mundo!")
	// go escrever("Programando em GO!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
