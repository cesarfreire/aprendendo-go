package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	//import implícito com _
)

func main() {
	// estrutura connectionstring usuario:senha@/banco
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		log.Fatal(erro)
	}

	// defer para fechar a conexão por último
	defer db.Close()

	// garante que a conexão foi feita com sucesso
	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão feita com sucesso.")

	// executa uma query
	linhas, erro := db.Query("select * from usuarios;")

	if erro != nil {
		log.Fatal(erro)
	}

	// fecha o cursor
	defer linhas.Close()

}
