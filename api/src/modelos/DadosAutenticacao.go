package modelos

//Contem o Token e ID do usuario autenticado
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
