package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Representa um repositorio de publicacoes
type Publicacoes struct {
	db *sql.DB
}

// Cria um repositorio de publicacoes
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Insere uma publicacao no banco de dados
func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// Buscar uma unica publicação do banco de dados pela ID dela
func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {
	linha, erro := repositorio.db.Query(`
		SELECT p.*, u.nick
		FROM publicacoes p INNER JOIN usuarios u
		ON u.id = p.autor_id
		WHERE p.id = ?
		`, publicacaoID,
	)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	defer linha.Close()

	var publicacao modelos.Publicacao
	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

// Traz as publicações do propior usuario e dos que ele segue
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
	select DISTINCT p.*, u.nick from publicacoes p
	inner join usuarios u on u.id = p.autor_id 
	inner join seguidores s on p.autor_id = s.usuario_id 
	WHERE u.id = ? or s.seguidor_id = ?
	order by 1 desc
	`, usuarioID, usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// altera os dados de uma publicação do db
func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao modelos.Publicacao) error {
	statement, erro := repositorio.db.Prepare("UPDATE publicacoes SET titulo = ?, conteudo = ? WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); erro != nil {
		return erro
	}

	return nil
}

// Exlui uma publicação do banco de dados
func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE FROM publicacoes WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}

// Traz todas as publicações de um usuario especifico
func (repositorio Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		SELECT p.*, u.nick FROM publicacoes p
		JOIN usuarios u on u.id = p.autor_id
		WHERE p.autor_id = ?`,
		usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Adiciona uma curtida na publicacao
func (repositorio Publicacoes) Curtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare("UPDATE publicacoes SET curtidas = curtidas + 1 WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}

// Subtrai uma curtida na publicacao
func (repositorio Publicacoes) Descurtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(`
		UPDATE publicacoes SET curtidas =
		CASE WHEN curtidas > 0 THEN curtidas -1
		ELSE curtidas END
		WHERE id = ?
	`)
	if erro != nil {
		return erro
	}
	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}
