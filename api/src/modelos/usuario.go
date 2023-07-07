package modelos

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	ID 				uint   		`json:"id,omitempty"`
	Nome 			string 		`json:"nome,omitempty"`
	Nick 			string 		`json:"nick,omitempty"`
	Senha 			string 		`json:"senha,omitempty"`
	CriadoEm 		time.Time 	`json:"criadoEm,omitempty"`
}


//FUNCAO QUE VAI CHAMAR OS METODOS PARA VALIDAR E FORMATAR O USUSARIO RECEBIDO
func(Usuario *Usuario) Preparar() error{
	if erro := Usuario.Validar(); erro != nil {
		return erro
	}
	Usuario.Formatar()
	return nil
}


func (usuario *Usuario) Validar() error  {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatorio e não pode estar em branco")
	}
	if usuario.Nick == "" {
		return errors.New("O nick é obrigatorio e não pode estar em branco")
	}
	if usuario.Senha == "" {
		return errors.New("O senha é obrigatorio e não pode estar em branco")
	}
	return nil
}


func (Usuario * Usuario) Formatar()  {
	Usuario.Nome = strings.TrimSpace(Usuario.Nome)
	Usuario.Nick = strings.TrimSpace(Usuario.Nick)
}
