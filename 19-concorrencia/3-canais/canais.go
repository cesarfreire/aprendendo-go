package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá mundo!", canal)

	fmt.Println("Depois da função começar a ser executada.")

	// recebendo um valor
	// for {
	// 	// recebe um bool se o canal ainda está aberto
	// 	mensagem, aberto := <-canal
	// 	//caso já ter sido fechado, finaliza o for
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	// refatorado
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		// mandando um valor pra dentro do canal
		canal <- texto
		time.Sleep(time.Second)
	}
	// fecha o canal
	close(canal)
}
