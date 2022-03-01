package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Connect abre a conexão com o banco de dados e a retorna
func Connect() (*sql.DB, error) {
	db, error := sql.Open("mysql", config.StringBDConnection)

	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		db.Close()

		return nil, error
	}

	return db, nil
}
