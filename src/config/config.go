package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringBDConnection é a string de conexão com o MySql
	StringBDConnection = ""

	// Port é a porta onde a API vai estar rodando
	Port = 0
)

// Load vai inicializar as variáveis de ambiente
func Load() {
	var error error

	if error = godotenv.Load(); error != nil {
		log.Fatal(error)
	}

	Port, error = strconv.Atoi(os.Getenv("PORT"))

	if error != nil {
		Port = 9000
	}

	StringBDConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))
}
