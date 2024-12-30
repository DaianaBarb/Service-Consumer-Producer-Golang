package postegres

import (
	"database/sql"
	"log"
	"os"
)

func Connect() *sql.DB {

	db, err := sql.Open("postegres", os.Getenv("DB_TENANT_1"))
	defer db.Close()
	if err != nil {
		log.Fatalf("connect DB error")
	}

	return db

}
