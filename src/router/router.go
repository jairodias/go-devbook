package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Gerar ai retornar um router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()

	return routes.Setting(r)
}
