package middlewares

import (
	"api/src/authenticate"
	"api/src/responses"
	"log"
	"net/http"
)

// Logger escrever informações da requisição como log no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Authenticate verifica se o usuario fazendo a requisição está logado
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if error := authenticate.ValidarToken(r); error != nil {
			responses.Error(w, http.StatusUnauthorized, error)
			return
		}
		next(w, r)
	}
}
