package modelos

//Contem o ID e Token do usuario Autenticado
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
