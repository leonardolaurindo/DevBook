package modelos

import "time"

//representa uma publicação feita por um usuario
type Publicacao struct {
	ID        uint64    `json:"id,omitempry"`
	Titulo    string    `json:"titulo,omitempry"`
	Conteudo  string    `json:"conteudo,omitempry"`
	AutorID   uint64    `json:"autorId,omitempry"`
	AutorNick string    `json:"autorNick,omitempry"`
	Curtidas  uint64    `json:"curtidas"`
	CriadaEm  time.Time `json:"criadaEm,omitempry"`
}
