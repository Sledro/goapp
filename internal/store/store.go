package store

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

// NewDatabase - Create a new database conection. Schema will be migrated if
// not found
func NewDatabase(username, password, host, port, database string) *sqlx.DB {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, database)
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
