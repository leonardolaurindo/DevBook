package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// Função Generica, recebe status code passado, insere no header e transforma os dados em JSON
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

// Retorna um erro em JSON
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json: "erro"`
	}{
		Erro: erro.Error(),
	})
}
