package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// Adiciona uma nova pulbicação no banco de dados
func CriarPublicacoes(w http.ResponseWriter, r *http.Request) {
	//Pegar usuario ID do token
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	// Ler o corpo da requisicao
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	//Criar publicação baseada no modelo
	var publicacao modelos.Pulbicacao
	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	//Inserir autor ID (usuarioID) de dentro do token na publicacao
	publicacao.AutorID = usuarioID
	// abrir conexão com DB
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, errors.New("Erro aqui 01"))
		return
	}
	defer db.Close()

	//Cria repositorio de publicacoes
	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacao.ID, erro = repositorio.Criar(publicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, errors.New("Erro aqui 02"))
		return
	}

	//Caso não der erro, retornamos a publicacao cmo titulo conteudo, autor id e id da publicacao
	respostas.JSON(w, http.StatusCreated, publicacao)
}

// Traz as publicacoes do feed do usuario
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {

}

// Traz uma unica publicaco
func BuscarPublicaco(w http.ResponseWriter, r *http.Request) {

}

// Altera os dados de uma publicacao
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {

}

// Exclue uma publicacao
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {

}
