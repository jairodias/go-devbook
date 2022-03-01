package controllers

import "net/http"

// Create insere um usuário no banco de dados
func Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário"))
}

// ListAll lista todos os usuários do banco
func ListAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando todos os usuarios"))
}

// SearchById busca um único usuário no banco de acordo com o id passado
func SearchById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário"))
}

// Update alterar as informações do id passado no route params
func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário"))
}

// Delete exclui o registro de um usuário dentro da base de dados
func Detele(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário"))
}
