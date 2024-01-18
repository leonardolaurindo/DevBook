package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Função para chamar a API e criar o usuario
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	//Pega o corpo da requisição
	r.ParseForm()
	//Map dos campos
	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		log.Fatal(erro)
	}

	response, erro := http.Post("http://localhost:5000/usuarios", "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		log.Fatal(erro)
	}
	defer response.Body.Close()

	fmt.Println(response.Body)

}
