package repositorio

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

//ESSE STRUCT RECEBE O BANCO
//A CONEXAO VAI SER ABERTA NO CONTROLLER E VAI SER PASSADA PARA CÁ E AQUI SERA FEITA A INTERACAO
//E DENTRO DESSE STRUCT QUE SERA FEITA OS METODOS QUE FARA A COMUNICACAO COM O BANCO DE DADOS
type Usuarios struct {
	db *sql.DB
}

//FUNCAO QUE CRIA REPOSITORIO DE USUARIOS
//FUNCAO VAI RECEBER UM BANCO QUE VAI SER ABERTA LA PELO CONTROLLER E O CONTROLLER VAI CHAMAR ESSA FUNCAO 
//E ESSA FUNCAO VAI PEGAR O BANCO E VAI JOGAR NO STRUCT USUARIOS
func NovoRepositorioDeUsuario(db *sql.DB) *Usuarios{
	return &Usuarios{db}
}

//METODOS QUE INSERE UM USUARIO NO BANCO DE DADOS
func (repositorio Usuarios) CriarUsuario(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuario (nome, nick, senha) values (?,?,?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()
	
	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Senha)
	if erro!= nil {
		return 0, erro
	}

	ultimaIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimaIDInserido), nil

}

//FUNCAO QUE TRAS TODOS OS USUARIOS QUE ATENDAM UM FILTRO DE NOME OU NICK
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOunick%
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, criadoEm from usuario where nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario
	
	if erro = linhas.Scan(
		&usuario.ID,
		&usuario.Nome,
		&usuario.Nick,
		&usuario.CriadoEm,

	); erro != nil {
		return nil, erro
	}

	usuarios = append(usuarios, usuario)

	}
	return usuarios, nil
}
