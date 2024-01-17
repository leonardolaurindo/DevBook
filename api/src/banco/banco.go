package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver SQl
)

// Conectar abre a conexão com banco de dados e a retorna
func Conectar() (*sql.DB, error) {
	//Abre a conexão
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}
	// Verifica se existe erro na conexão, fecha e retorna
	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}
	// Retorna conexão
	return db, nil
}
