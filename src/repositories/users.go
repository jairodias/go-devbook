package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

// Search traz todos os usuários que atendem um filtro de nome ou nick
func (repository users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	lines, error := repository.db.Query(
		"SELECT id, name, nick, email, created_at FROM users WHERE name LIKE ? OR nick LIKE ?",
		nameOrNick,
		nameOrNick,
	)

	if error != nil {
		return nil, error
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if error = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.Created_At,
		); error != nil {
			return nil, error
		}

		users = append(users, user)
	}

	return users, nil
}

// FindById busca um usuário de acordo com o id informado
func (repository users) FindById(ID uint64) (models.User, error) {
	line, error := repository.db.Query(
		"SELECT id, name, nick, email, created_at FROM users WHERE id = ?",
		ID,
	)
	if error != nil {
		return models.User{}, error
	}
	defer line.Close()

	var user models.User
	if line.Next() {
		if error = line.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.Created_At,
		); error != nil {
			return models.User{}, error
		}
	}

	return user, nil
}

func (repository users) Update(ID uint64, user models.User) error {
	statement, error := repository.db.Prepare(
		"UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?",
	)

	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(user.Name, user.Nick, user.Email, ID); error != nil {
		return error
	}

	return nil
}

// Delete exclui as informaçoes de um usuario no banco de dados
func (repository users) Delete(ID uint64) error {
	statement, error := repository.db.Prepare("DELETE FROM users WHERE id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(ID); error != nil {
		return error
	}

	return nil
}

// FindByEmail busca um usuário por email e retorna o seu ID e senha com hash
func (repository users) FindByEmail(email string) (models.User, error) {
	line, error := repository.db.Query("SELECT id, password FROM users WHERE email = ?", email)

	if error != nil {
		return models.User{}, error
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if error = line.Scan(&user.ID, &user.Password); error != nil {
			return models.User{}, error
		}
	}

	return user, nil
}
