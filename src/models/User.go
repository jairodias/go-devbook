package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User representa um usuário utilizando a rede social
type User struct {
	ID         uint64    `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Nick       string    `json:"nick,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	Created_At time.Time `json:"created_at,omitempty"`
}

// Prepare vai chamar os métodos para validar e formatar os dados do usuário recebido
func (user *User) Prepare(etapa string) error {
	if error := user.validate(etapa); error != nil {
		return error
	}

	if error := user.format(etapa); error != nil {
		return error
	}

	return nil
}

func (user *User) validate(etapa string) error {
	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Nick == "" {
		return errors.New("nick is required")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if error := checkmail.ValidateFormat(user.Email); error != nil {
		return error
	}

	if etapa == "create" && user.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func (user *User) format(etapa string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if etapa == "create" {
		passwordHashed, error := security.Hash(user.Password)
		if error != nil {
			return error
		}

		user.Password = string(passwordHashed)
	}

	return nil
}
