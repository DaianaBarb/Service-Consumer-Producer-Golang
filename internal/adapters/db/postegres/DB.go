package postegres

import (
	"database/sql"
	"log"
	"os"
)

func Connect() *sql.DB{
	db, err := sql.Open("postegres", os.Getenv("DB_TENANT_1"))

  if err!= nil{
	log.Fatalf("connect DB error")
  }
  return db
}