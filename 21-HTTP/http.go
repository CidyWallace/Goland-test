package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo"))
}

func usuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregando página de usuários"))
}

func main() {
	//HTTP É UM PROTOCOLO DE COMUNICAÇÃO, BASE DE COMUNICAÇÃO WEB

	//CLIENTE (FAZ REQUISIÇÕES) - SERVIDOR (PROCESSA REQUISIÇÕES E ENVIA RESPOSTA)

	//Request - Response

	//Rotas
	//URI - Identificador do Recurso
	//Método - GET, POST, PUT, DELETE

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuario", usuario)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
