package requisicoes

import (
	"io"
	"net/http"
	"webapp/src/cookies"
)

// Coloca o token na requisição
func FazerRequisicaoComAutenticacao(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
	//Cria um requisição que sera mandada
	request, erro := http.NewRequest(metodo, url, dados)
	if erro != nil {
		return nil, erro
	}
	//Precisa de um cliente http para ser feito a requisição, e só sera feita com o cookie já no header
	//Ler o Cookie ignorando o erro
	cookie, _ := cookies.Ler(r)
	//Passa os valores para o header da requisição (request)
	request.Header.Add("Authorization", "Bearer "+cookie["token"])
	//Cria o client HTTp para fazer a requisição
	client := &http.Client{}
	//recebe a requisição que foi criada (request)
	response, erro := client.Do(request)
	if erro != nil {
		return nil, erro
	}

	return response, nil
}
