package repositories

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

// UserRepository cria um repositório de usuários
func UserRepository(db *sql.DB) *users {
	return &users{db}
}

// Create insere um usuário no banco de dados [referencia do struct]
func (repository users) Create(user models.User) (uint64, error) {
	statement, error := repository.db.Prepare("INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)")

	if error != nil {
		return 0, error
	}

	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.Password)

	if error != nil {
		return 0, error
	}

	lastInsertId, error := result.LastInsertId()

	if error != nil {
		return 0, error
	}

	return uint64(lastInsertId), nil
}
