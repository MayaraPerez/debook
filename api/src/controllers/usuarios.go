package controllers

//TODO MUNDO QUE ESTA DENTRO DESSE PACOTE CONTROLLLER VAI LIDAR COM AS FUNCOES HTTP
//VAI RECEBER AS REQUISICOES E VAI DEVOLVER AS RESPOSTAS

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorio"
	"api/src/respostas"
	"encoding/json"
	"fmt"
	"io/ioutil" // retorna dois valores
	"net/http"
	"strings"
)

//FUNCAO PARA CRIAR USUARIO QUE ESTA VINDO LA DO PACOTE MODELOS
func CriarUsuario(w http.ResponseWriter, r *http.Request){
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil{
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	//CRIAMOS A CONEXAO COM BANCO DE DADOS
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	
	//PASSAMOS ESSA CONEXAO PARA O REPOSITORIO VAI SER A CAMADA QUE VAI INTERAGIR COM O BANCO DE DADOS
	//E ATRAVES DOS METODOS QUE TEM NESSE REPOSITORIO QUE JA ESTA COM A CONEXAO COM BANCO VAMOS FAZER AS INSERCOES 
	repositorio := repositorio.NovoRepositorioDeUsuario(db)
	usuarioID, erro := repositorio.CriarUsuario(usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	w.Write([]byte(fmt.Sprintf("ID inserido %d", usuarioID)))
}
//FUNCAO PARA BUSCAR TODOS USUARIO
func BuscandoUsuario(w http.ResponseWriter, r *http.Request){
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorios := repositorio.NovoRepositorioDeUsuario(db)
	usuarios, erro := repositorios.Buscar(nomeOuNick)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusOK, usuarios)
}



//FUNCAO PARA BUSCAR ID USUARIO
func BuscandoUsuarioId(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscando Usuario id"))
}
//FUNCAO PARA ATUALIZAR USUARIO
func AtualizandoUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Atualizando Usuario"))
}
//FUNCAO PARA DELETAR USUARIO
func DeletandoUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Deletando Usuario"))
}


