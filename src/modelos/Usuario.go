package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario representa um usuario na API
type Usuario struct {
	ID       uint64    `json :"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

// Chama os metodos para validar e formatar o usuario rececbido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O Nome é obrigatorio e não pode estar em branco!")
	}
	if usuario.Nick == "" {
		return errors.New("O Nick é obrigatorio e não pode estar em branco!")
	}
	if usuario.Email == "" {
		return errors.New("O Email é obrigatorio e não pode estar em branco!")
	}
	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O email é invalidoo")
	}
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A Senha é obrigatorio e não pode estar em branco!")
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}
		usuario.Senha = string(senhaComHash)
	}

	return nil

}
