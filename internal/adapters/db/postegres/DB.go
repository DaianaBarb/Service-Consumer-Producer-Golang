package postegres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	uri := os.Getenv("DB_TENANT_1")

	db, err := sql.Open("postgres", uri)

	if err != nil {
		log.Fatalf("connect DB error")
	}

	fmt.Println("Conexão bem-sucedida com o PostgreSQL!")
	return db

}
