package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá usuários!"))
}

func main() {
	// http é um protocolo de comunicação
	//base da comunicação web

	// cliente (faz requisição) - servidor (processa e responde)
	// request - response

	// rotas
	// URI - identificador do recurso
	// Método - Ação que irá fazer com o recurso - GET, POST, DELETE, PUT

	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
