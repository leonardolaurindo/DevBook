package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// Erro representa a resposta de Erro da API
type ErroAPI struct {
	Erro string `json:"erro"`
}

// Retorna uma respota em formado JSOn para requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}

}

func TratarStatusCodeDeErro(w http.ResponseWriter, r *http.Response) {
	var erro ErroAPI
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}
