package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	//não é necessário usar o := nesse caso
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma, sub := calculosMatematicos(5, 6)
	fmt.Println(soma, sub)
}
