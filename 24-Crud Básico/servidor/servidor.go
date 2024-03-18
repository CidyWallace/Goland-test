package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuario insere um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler corpo da requisição"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(body, &usuario); erro != nil {
		w.Write([]byte("Erro ao Converter o usuário para struct"))
	}

	//Abrindo conexão com o banco
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
	}
	defer db.Close()

	//PREPRARE STATEMENT
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
	}
	defer statement.Close()

	insert, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement!"))
	}

	idInserido, erro := insert.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter ID inserido!"))
	}

	//STATUS CODES
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso: Id %d", idInserido)))
}

// BuscarUsuarios traz todos os usuarios salvos no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
	}
	defer db.Close()

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar os usuários!"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario

	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário!"))
			return
		}

		usuarios = append(usuarios, usuario)
	}
	w.WriteHeader(http.StatusOK)

	//Responde Body
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter os usuários para json"))
		return
	}
}

// BuscarUsuario traz um usuarios específico salvo no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametro["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar usuário"))
		return
	}
	defer linha.Close()

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao buscar o usuário"))
		}
	}
	w.WriteHeader(http.StatusOK)

	//Response Body
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter para json"))
	}
}

// AtualizarUsuario altera os dados de um usuário já salvo no bando de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(body, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeletarUsuario remove um usuário do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	paramentros := mux.Vars(r)
	ID, erro := strconv.ParseUint(paramentros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statemente, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statemente"))
		return
	}
	defer statemente.Close()

	if _, erro := statemente.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar o usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
