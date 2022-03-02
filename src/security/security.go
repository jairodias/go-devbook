package security

import "golang.org/x/crypto/bcrypt"

// Hash recebe uma string e coloca um nela nela
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Verify compara uma senha e um hash e retorna se elas são iguais
func Verify(passowordHashed, passwordString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passowordHashed), []byte(passwordString))
}
