package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	// tem que fazer essa parada pra não ficar aparentemente "hardcoded" e não vazar nada sobre o banco
	connectionString := fmt.Sprintf("host%s port=%s user=%s password=%s dbname=%s sslmode=disable connection_timeout=5",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("não foi possivel conectar ao banco de dados: ", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("não foi possivel conectar ao banco de dados: ", err)
	}

	file, err := os.ReadFile("internal/db/migrations/setup.sql")
	if err != nil {
		log.Fatal("não foi possivel ler o arquivo SQL: ", err)
	}
	_, err = db.Exec(string(file))
	if err != nil {
		log.Fatal("não foi possivel executar o arquivo SQL: ", err)
	}
	fmt.Println("conectado ao banco de dados")

	return db, nil
}
