package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {
	// chave := os.Getenv("DATA_KEY")
	conexao := fmt.Sprintf(
		"user=postgres dbname=storage_go password=%s host=localhost sslmode=disable",
		os.Getenv("DATA_KEY"))
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
