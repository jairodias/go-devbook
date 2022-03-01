package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Create insere um usuário no banco de dados
func Create(w http.ResponseWriter, r *http.Request) {
	body, error := ioutil.ReadAll(r.Body)

	if error != nil {
		log.Fatal(error)
	}

	var user models.User
	if error = json.Unmarshal(body, &user); error != nil {
		log.Fatal(error)
	}

	db, error := database.Connect()
	if error != nil {
		log.Fatal(error)
	}

	repository := repositories.UserRepository(db)
	repository.Create(user)
}

// ListAll lista todos os usuários do banco
func ListAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando todos os usuarios"))
}

// SearchById busca um único usuário no banco de acordo com o id passado
func SearchById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuário"))
}

// Update alterar as informações do id passado no route params
func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Alterando usuário"))
}

// Delete exclui o registro de um usuário dentro da base de dados
func Detele(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário"))
}
