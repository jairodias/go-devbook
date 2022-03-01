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
func (u users) Create(user models.User) (uint64, error) {
	return 0, nil
}
