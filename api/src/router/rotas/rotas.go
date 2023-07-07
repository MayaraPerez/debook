package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rotas struct {
	Url 				string
	Metodo  			string
	Funcao  			func(http.ResponseWriter, *http.Request)
	ResquerAutenticador bool
}

func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios

	for _, rota := range rotas {
		r.HandleFunc(rota.Url, rota.Funcao).Methods(rota.Metodo)
	}
	return r
}
