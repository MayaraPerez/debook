package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rotas {
	{
		Url: 				"/usuarios",
		Metodo: 			http.MethodPost,
		Funcao: 			controllers.CriarUsuario,
		ResquerAutenticador: false,
	},
	{
		Url: 				"/usuarios",
		Metodo:				 http.MethodGet,
		Funcao: 			 controllers.BuscandoUsuario,
		ResquerAutenticador: false,
	},
	{
		Url: 				"/usuarios/{usuarioId}",
		Metodo: 			http.MethodGet,
		Funcao: 			controllers.BuscandoUsuarioId,
		ResquerAutenticador: false,
	},
	{
		Url: 				"/usuarios/{usuarioId}",
		Metodo:				 http.MethodPut,
		Funcao: 			controllers.AtualizandoUsuario,
		ResquerAutenticador: false,
	},
	{
		Url: 				"/usuarios/{usuarioId}",
		Metodo: 			http.MethodDelete,
		Funcao: 			controllers.DeletandoUsuario,
		ResquerAutenticador: false,
	},
}