package main

import (
	"bytes"
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
	// struct para json
	c := cachorro{"Rex", "Dalmata", 3}

	cachorroEmJSON, erro := json.Marshal(c)

	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorroEmJSON) // []uint8 - slice de bytes

	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	// map para json
	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro2EmJSON)
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))

}
