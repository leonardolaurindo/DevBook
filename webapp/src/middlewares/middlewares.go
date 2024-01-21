package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Escreve infomações da requisição no terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

// Verifica a existencia de cookies
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	//Antes de executar a proma função, verifica se os dados estão no cookie
	return func(w http.ResponseWriter, r *http.Request) {
		if _, erro := cookies.Ler(r); erro != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}
		proximaFuncao(w, r)
	}
}
