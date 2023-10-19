package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	//json para struct
	cachorroEmJSON := `{"nome":"Rex","raca":"Dalmata","idade":3}`

	var c cachorro

	// & para endereco de memória do c
	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	//json para map
	cachorro2EmJSON := `{"nome":"Bituca","raca":"Pintcher"}`
	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)
}
