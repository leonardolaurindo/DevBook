package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Representa um repositirio de usuarios
type usuarios struct {
	db *sql.DB
}

// Cria um repositorio de usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// Criar insere um usuario no banco de dados
func (repositorio usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	// Insere banco dentro do repositorio
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO usuarios (nome, nick, email, senha) VALUES (?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()
	// Envia dados do usuario que chegaram por parametro
	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}
