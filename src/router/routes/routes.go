package routes

import "net/http"

// Rota representa todas as rotas da API
type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}
