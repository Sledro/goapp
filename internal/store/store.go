package store

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

// NewDatabase - Create a new database conection. Schema will be migrated if
// not found
func NewDatabase(username, password, host, port, database string) *sqlx.DB {
	connString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", username, password, host, port, database)
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		log.Fatalln(err)
	}

	// OR: Read migrations from a folder:
	migrations := &migrate.FileMigrationSource{
		Dir: "schema",
	}

	n, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
	return db
}
