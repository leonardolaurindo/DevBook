package cookies

import (
	"net/http"
	"time"
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

// Retorna os valores armazenados no cookie (map)
func Ler(r *http.Request) (map[string]string, error) {
	//Pega os cookies codificados
	cookie, erro := r.Cookie("dados")
	if erro != nil {
		return nil, erro
	}

	valores := make(map[string]string)
	// Decodifica os dados e joga dentro do map valores.
	if erro = s.Decode("dados", cookie.Value, &valores); erro != nil {
		return nil, erro
	}

	return valores, nil
}

// Remove os valores armazenados no cookie
func Deletar(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})
}
