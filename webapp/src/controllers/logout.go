package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

// Remove os dados de autenticação slavos no navegador do usuario
func FazerLogout(w http.ResponseWriter, r *http.Request) {
	cookies.Deletar(w)
	http.Redirect(w, r, "/login", 302)
}
