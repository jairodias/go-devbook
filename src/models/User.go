package models

import (
	"errors"
	"strings"
	"time"
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
func (user *User) Prepare() error {
	if error := user.validate(); error != nil {
		return error
	}

	user.format()
	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Nick == "" {
		return errors.New("nick is required")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if user.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
