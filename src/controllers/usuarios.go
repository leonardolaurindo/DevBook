package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Controller, quem lida com as funções HTTP
// CriarUsuario insere usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		fmt.Println("Erro aqui 1")
		log.Fatal(erro)
	}
	// Criando usuario do pacote modelos.
	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		fmt.Println("Erro aqui 2")
		log.Fatal(erro)
	}
	// Cria uma conexão com o banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		fmt.Println("Erro aqui 3")
		log.Fatal(erro)
	}
	// Passa a conexão para o repositorio
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioID, erro := repositorio.Criar(usuario)
	if erro != nil {
		fmt.Println("Erro aqui 4")
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("ID inserido: %d", usuarioID)))
}

// BuscarUsuarios busca todos os usuarios salvo no banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os Usuarios!"))
}

// BuscarUsuario busca um usuario salvo no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um Usuario!"))
}

// AtualizarUsuario altera as informaç~çoes de um usuario no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuario!"))
}

// DeletarUsuario deleta um usuario do banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuario!"))
}
