package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
)

// Chama a API para cadastrar uma publicação no banco de dados
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//Criar publicação com um json mashal
	publicacao, erro := json.Marshal(map[string]string{
		"titulo":   r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//caso não der erro, cria uma variavel que é a URL da API que vamos fazer a criação da publicação
	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	//Faze a requisição
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(publicacao))
	//Faz tratamento de erro
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	//Devolve uma resposta para o FrontEnd
	respostas.JSON(w, response.StatusCode, nil)

}
