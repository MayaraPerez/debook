package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

//FUNCAO QUE RECEBE UM STATUSCODE QUE VAI SER PASSADO VAI COLOCAR ESSE STATUSCODE E COLOCA NO WRITEHEADER
// E DEPOIS VAI PEGAR OS DADOS E VAI PASSAR PARA JSON
func JSON(w http.ResponseWriter, statusCode int, dados interface{})  {
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w) .Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

func Erro(w http.ResponseWriter, statusCode int, erro error)  {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}