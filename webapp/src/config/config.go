package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//Reporesenta a URL para comunicação com a API
	APIURL = ""
	//Porta onde a aplicação web esta rodando
	Porta = 0
	//Utilizada para autenticar o cooki
	HashKey []byte
	//Para criptografar os dados do cookie
	BlockKey []byte
)

// Inicializa as variaves de ambiente
func Carregar() {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("APPPORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	APIURL = os.Getenv("APIURL")
	HashKey = []byte(os.Getenv("HASHKEY"))
	BlockKey = []byte(os.Getenv("BLOCKKEY"))

}
