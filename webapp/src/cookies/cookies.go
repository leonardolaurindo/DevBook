package cookies

import (
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// Ultiliza as variaveis de ambiente para a criação do SecureCookie
func Configurar() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// Regista as informações de autenticação
func Salvar(w http.ResponseWriter, ID, token string) error {
	//Cria os dados a partir de um map
	dados := map[string]string{
		"id":    ID,
		"token": token,
	}
	//Codifica os dados
	dadosCodificados, erro := s.Encode("dados", dados)
	if erro != nil {
		return erro
	}

	// Inserindo cookie no browser
	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    dadosCodificados,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}
